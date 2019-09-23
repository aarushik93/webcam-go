package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

const dir = "."

func main() {
	fs := http.FileServer(http.Dir(dir))
	log.Printf("Serving "+dir+" on http://localhost:%+v", os.Getenv("PORT"))

	http.ListenAndServe(":"+os.Getenv("PORT"), http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Cache-Control", "no-cache")
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}
		fs.ServeHTTP(resp, req)
	}))
}

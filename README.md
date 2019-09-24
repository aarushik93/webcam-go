This is a photobooth filter application implemented in JS and Go. The filtering is all done in Go and compiled down to WASM. 
Visit the working app at: https://wasm-webcam.herokuapp.com


# Compiling
Requires Go 1.13 or above.

GOOS=js GOARCH=wasm go build -o main.wasm main.go


# Testing
In browser testing, so you don't need to set up node js. Follow:  https://github.com/agnivade/wasmbrowsertest


# Docker
`docker build -t wasm-go .`

`docker run --rm -it -p 8080:8080 wasm-go`

# Contact details
twitter.com/aarushikansal

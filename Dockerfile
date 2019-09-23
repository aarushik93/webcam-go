FROM golang:1.13 AS builder
ENV GOOS=js GOARCH=wasm
COPY . $GOPATH/src/github.com/aarushik93/wasm-webcam/
WORKDIR  $GOPATH/src/github.com/aarushik93/wasm-webcam/

ENV PORT=8080
EXPOSE 8080

RUN GOOS=js GOARCH=wasm go build -o $GOPATH/src/github.com/aarushik93/wasm-webcam/main.wasm $GOPATH/src/github.com/aarushik93/wasm-webcam/main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH= go build $GOPATH/src/github.com/aarushik93/wasm-webcam/server/serve.go

CMD ["./serve"]




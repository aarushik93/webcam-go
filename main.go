package main

import (
	"github.com/aarushik93/wasm-webcam/filter"
	"syscall/js"
)

var uint8Array = js.Global().Get("Uint8Array")
var threshold = 127

func applyFilter(this js.Value, args []js.Value) interface{} {

	val := args[0].Int()
	length := args[1].Length()

	jsPixels := make(filter.Pixels, length)

	_ = js.CopyBytesToGo(jsPixels, args[1])

	if val == 0 {
		jsPixels.MakeGrey()

	} else if val == 1 {
		jsPixels.Invert()

	} else if val == 2 {
		jsPixels.MakeNoise()

	} else if val == 3 {
		jsPixels.MakeRed()
	}

	buf := uint8Array.New(len(jsPixels))
	js.CopyBytesToJS(buf, jsPixels)

	return buf
}


func main() {
	done := make(chan struct{}, 0)

	applyFilterFunc := js.FuncOf(applyFilter) // wrapper function
	js.Global().Set("applyFilter", applyFilterFunc)

	defer applyFilterFunc.Release()
	<-done
}

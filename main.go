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
		//data = makeInvert(args[1])

	} else if val == 2 {
		jsPixels.MakeNoise()
		//data = noise(args[1])

	} else if val == 3 {
		jsPixels.MakeRed()
		//data = makeRed(args[1])

	}

	buf := uint8Array.New(len(jsPixels))
	js.CopyBytesToJS(buf, jsPixels)

	return buf
}

//func makegrey(arg js.Value) []byte {
//	length := arg.Length()
//	nongreyed := make(filter.Pixels, length)
//
//	_ = js.CopyBytesToGo(nongreyed, arg)
//	nongreyed.MakeGrey()
//
//	return nongreyed
//}
//
//func makeInvert(arg js.Value) []byte {
//	length := arg.Length()
//	nontint := make(filter.Pixels, length)
//
//	_ = js.CopyBytesToGo(nontint, arg)
//
//	nontint.Invert()
//
//	return nontint
//
//}
//
//func noise(arg js.Value) []byte {
//	length := arg.Length()
//	nontint := make(filter.Pixels, length)
//
//	_ = js.CopyBytesToGo(nontint, arg)
//	nontint.MakeNoise()
//	return nontint
//
//}
//
//func makeRed(arg js.Value) []byte {
//	length := arg.Length()
//	nontint := make(filter.Pixels, length)
//
//	_ = js.CopyBytesToGo(nontint, arg)
//	nontint.MakeRed()
//
//	return nontint
//
//}
//
//func doNothing(arg js.Value) []byte {
//	length := arg.Length()
//	ogPixels := make([]byte, length)
//	_ = js.CopyBytesToGo(ogPixels, arg)
//
//	return ogPixels
//}


func main() {
	done := make(chan struct{}, 0)

	applyFilterFunc := js.FuncOf(applyFilter) // wrapper function
	js.Global().Set("applyFilter", applyFilterFunc)

	defer applyFilterFunc.Release()
	<-done
}

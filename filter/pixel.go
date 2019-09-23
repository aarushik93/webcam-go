package filter

import "math/rand"

// Pixels is a byte array to represent image pixels.
type Pixels []byte

// MakeGrey applies a grey filter to pixels.
func (p Pixels) MakeGrey() {
	for i := 3; i < len(p); i += 4 {
		red := p[i-3]   //uint8(arg.Index(i-3).Int())
		green := p[i-2] //uint8(arg.Index(i-2).Int())
		blue := p[i-1]  //uint8(arg.Index(i-1).Int())

		grey := uint8(float64(red)*0.299 + float64(green)*0.587 + float64(blue)*0.114)
		p[i-3] = grey
		p[i-2] = grey
		p[i-1] = grey
	}
}

// Invert applies a invert filter to pixels.
func (p Pixels) Invert() {
	for i := 3; i < len(p); i += 4 {

		p[i-3] = 255 - p[i-3]
		p[i-2] = 255 - p[i-2]
		p[i-1] = 255 - p[i-1]
	}
}

// MakeNoise applies a random number to pixels.
func (p Pixels) MakeNoise() {
	random := 1 - rand.Float64() + 3
	for i := 3; i < len(p); i += 4 {

		p[i-3] = uint8(float64(p[i-3])*random)
		p[i-2] = uint8(float64(p[i-2])*random)
		p[i-1] = uint8(float64(p[i-1])*random)
	}
}

// MakeRed applies a red filter to pixels by making blue and green 0.
func (p Pixels) MakeRed() {
	for i := 3; i < len(p); i += 4 {
		p[i-2] = 0
		p[i-1] = 0
	}
}

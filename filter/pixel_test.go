package filter_test

import (
	"github.com/aarushik93/wasm-webcam/filter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeGrey(t *testing.T) {
	pixels := make(filter.Pixels, 4, 4)
	expected := make(filter.Pixels, 4, 4)

	pixels[0] = 150
	pixels[3] = 20

	expected[0] = 44
	expected[1] = 44
	expected[2] = 44
	expected[3] = 20

	pixels.MakeGrey()

	assert.Equal(t, expected, pixels)
}

func TestInvert(t *testing.T) {
	pixels := make(filter.Pixels, 4, 4)
	expected := make(filter.Pixels, 4, 4)

	pixels[0] = 150
	pixels[1] = 100
	pixels[2] = 50
	pixels[3] = 20

	expected[0] = 105
	expected[1] = 155
	expected[2] = 205
	expected[3] = 20

	pixels.Invert()

	assert.Equal(t, expected, pixels)
}

func TestNoise(t *testing.T) {
	pixels := make(filter.Pixels, 4, 4)

	pixels[0] = 150
	pixels[1] = 100
	pixels[2] = 50
	pixels[3] = 20


	pixels.MakeNoise()

	assert.Equal(t, uint8(20), pixels[3])
	assert.NotEqual(t, 50, pixels[2])
	assert.NotEqual(t, 150, pixels[1])
	assert.NotEqual(t, 150, pixels[0])
}

func TestRed(t *testing.T) {
	pixels := make(filter.Pixels, 4, 4)
	expected := make(filter.Pixels, 4, 4)

	pixels[0] = 150
	pixels[1] = 100
	pixels[2] = 50
	pixels[3] = 20

	expected[0] = 150
	expected[1] = 0
	expected[2] = 0
	expected[3] = 20

	pixels.MakeRed()

	assert.Equal(t, expected, pixels)
}

package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringShouldDisplayColorWhen000(t *testing.T) {
	expected := `hex: #000000
rgb: 0, 0, 0
hsl: 0°, 0%, 0%
`
	color, err := NewColorFromHexCode("000")
	assert.Nil(t, err)
	assert.Equal(t, expected, color.String())
}

func TestStringShouldDisplayColor(t *testing.T) {
	expected := `hex: #aabbcc
rgb: 170, 187, 204
hsl: 210°, 53%, 73%
`

	color, err := NewColorFromHexCode("abc")
	assert.Nil(t, err)
	assert.Equal(t, expected, color.String())
}

func TestFooShouldReturnError(t *testing.T) {
	_, err := NewColorFromHexCode("foo")
	assert.NotNil(t, err)
}

func Test000ShouldReturn000000(t *testing.T) {
	color, err := NewColorFromHexCode("000")
	assert.Nil(t, err)
	assert.Equal(t, "#000000", color.HexCode)
}

func Test000000ShouldReturn000000(t *testing.T) {
	color, err := NewColorFromHexCode("000000")
	assert.Nil(t, err)
	assert.Equal(t, "#000000", color.HexCode)
}

func TestHash000000ShouldReturn000000(t *testing.T) {
	color, err := NewColorFromHexCode("#000000")
	assert.Nil(t, err)
	assert.Equal(t, "#000000", color.HexCode)
}

func TestAbcShouldReturnAabbcc(t *testing.T) {
	color, err := NewColorFromHexCode("abc")
	assert.Nil(t, err)
	assert.Equal(t, "#aabbcc", color.HexCode)
}

func TestAabbccShouldReturnAabbcc(t *testing.T) {
	color, err := NewColorFromHexCode("aabbcc")
	assert.Nil(t, err)
	assert.Equal(t, "#aabbcc", color.HexCode)
}

func TestHashAabbccShouldReturnAabbcc(t *testing.T) {
	color, err := NewColorFromHexCode("#aabbcc")
	assert.Nil(t, err)
	assert.Equal(t, "#aabbcc", color.HexCode)
}

func TestABCShouldReturnAabbcc(t *testing.T) {
	color, err := NewColorFromHexCode("ABC")
	assert.Nil(t, err)
	assert.Equal(t, "#aabbcc", color.HexCode)
}

func TestAABBCCShouldReturnAabbcc(t *testing.T) {
	color, err := NewColorFromHexCode("AABBCC")
	assert.Nil(t, err)
	assert.Equal(t, "#aabbcc", color.HexCode)
}

func TestHashAABBCCShouldReturnAabbcc(t *testing.T) {
	color, err := NewColorFromHexCode("#AABBCC")
	assert.Nil(t, err)
	assert.Equal(t, "#aabbcc", color.HexCode)
}

func Test1a2b3cShouldReturn1a2b3c(t *testing.T) {
	color, err := NewColorFromHexCode("1a2b3c")
	assert.Nil(t, err)
	assert.Equal(t, "#1a2b3c", color.HexCode)
}

func TestHash1a2b3cShouldReturn1a2b3c(t *testing.T) {
	color, err := NewColorFromHexCode("#1a2b3c")
	assert.Nil(t, err)
	assert.Equal(t, "#1a2b3c", color.HexCode)
}

func Test000ShouldReturnRGB(t *testing.T) {
	color, err := NewColorFromHexCode("000")
	assert.Nil(t, err)
	assert.Equal(t, 0, color.Red)
	assert.Equal(t, 0, color.Green)
	assert.Equal(t, 0, color.Blue)
}

func TestFffShouldReturnRGB(t *testing.T) {
	color, err := NewColorFromHexCode("fff")
	assert.Nil(t, err)
	assert.Equal(t, 255, color.Red)
	assert.Equal(t, 255, color.Green)
	assert.Equal(t, 255, color.Blue)
}

func TestAbc123ShouldReturnRGB(t *testing.T) {
	color, err := NewColorFromHexCode("abc123")
	assert.Nil(t, err)
	assert.Equal(t, 171, color.Red)
	assert.Equal(t, 193, color.Green)
	assert.Equal(t, 35, color.Blue)
}

func Test000ShouldReturnHSL(t *testing.T) {
	color, err := NewColorFromHexCode("000")
	assert.Nil(t, err)
	assert.Equal(t, 0, color.Hue)
	assert.Equal(t, 0, color.Saturation)
	assert.Equal(t, 0, color.Lightness)
}

func TestFffShouldReturnHSL(t *testing.T) {
	color, err := NewColorFromHexCode("fff")
	assert.Nil(t, err)
	assert.Equal(t, 0, color.Hue)
	assert.Equal(t, 0, color.Saturation)
	assert.Equal(t, 100, color.Lightness)
}

func TestAbc123ShouldReturnHSL(t *testing.T) {
	color, err := NewColorFromHexCode("abc123")
	assert.Nil(t, err)
	assert.Equal(t, 68, color.Hue)
	assert.Equal(t, 69, color.Saturation)
	assert.Equal(t, 44, color.Lightness)
}

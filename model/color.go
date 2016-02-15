package model

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("^#?[0-9a-fA-F]{3}([0-9a-fA-F]{3})?$")

type Color struct {
	HexCode    string
	Red        int
	Green      int
	Blue       int
	Hue        int
	Saturation int
	Lightness  int
}

func NewColorFromHexCode(hexCode string) (*Color, error) {
	if !re.Match([]byte(hexCode)) {
		return nil, errors.New("You must specify a valid hex color")
	}

	hexCode = normalizeHexCode(hexCode)
	red, green, blue, err := calculateRGBFromHex(hexCode)
	if err != nil {
		return nil, err
	}
	hue, saturation, lightness := calculateHSLFromRGB(red, green, blue)

	return &Color{
		HexCode:    hexCode,
		Red:        red,
		Green:      green,
		Blue:       blue,
		Hue:        hue,
		Saturation: saturation,
		Lightness:  lightness,
	}, nil
}

func (color *Color) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("hex: %s\n", color.HexCode))
	buffer.WriteString(fmt.Sprintf("rgb: %d, %d, %d\n", color.Red, color.Green, color.Blue))
	buffer.WriteString(fmt.Sprintf("hsl: %d°, %d%%, %d%%\n", color.Hue, color.Saturation, color.Lightness))

	return buffer.String()
}

func normalizeHexCode(hexCode string) string {
	// Cut the leading # and lowercase it
	hexCode = strings.TrimLeft(hexCode, "#")
	hexCode = strings.ToLower(hexCode)

	// Normalize it to #xxxxxx
	var buffer bytes.Buffer
	buffer.WriteString("#")

	if len(hexCode) == 6 {
		buffer.WriteString(hexCode)
	} else {
		for _, b := range []byte(hexCode) {
			buffer.WriteByte(b)
			buffer.WriteByte(b)
		}
	}

	return buffer.String()
}

func calculateRGBFromHex(hexCode string) (int, int, int, error) {
	var red, green, blue []byte
	red, err := hex.DecodeString(hexCode[1:3])
	if err == nil {
		green, err = hex.DecodeString(hexCode[3:5])
	}
	if err == nil {
		blue, err = hex.DecodeString(hexCode[5:7])
	}
	return int(red[0]), int(green[0]), int(blue[0]), err
}

func calculateHSLFromRGB(red int, green int, blue int) (int, int, int) {
	r := float64(red) / 255.0
	g := float64(green) / 255.0
	b := float64(blue) / 255.0

	min := math.Min(r, math.Min(g, b))
	max := math.Max(r, math.Max(g, b))
	span := max - min

	hue, saturation, lightness := 0, 0, int((max+min)/0.02)

	if span != 0.0 {
		if lightness < 50 {
			saturation = int(100 * span / (max + min))
		} else {
			saturation = int(100 * (2 - max - min))
		}

		epsilon := 0.0000000001
		if math.Abs(r-max) < epsilon {
			hue = int(60 * math.Mod(((g-b)/span), 6))
		} else if math.Abs(g-max) < epsilon {
			hue = int(60 * (((b - r) / span) + 2))
		} else {
			hue = int(60 * (((r - g) / span) + 4))
		}
	}
	if hue < 0 {
		hue += 360
	}
	return hue, saturation, lightness
}

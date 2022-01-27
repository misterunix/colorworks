package hsl

import (
	"math"
	"testing"
)

func TestRGBtoHSL(t *testing.T) {
	r := uint8(24)
	g := uint8(98)
	b := uint8(118)
	h, s, l := RGBtoHSL(r, g, b)

	h1 := math.Round(h)
	h2 := int(h1)
	if h2 != 193 {
		t.Errorf("hue = %d, want %d", h2, 193)
	}

	s1 := s * 100.0
	s2 := int(math.Round(s1))

	if s2 != 66 {
		t.Errorf("sat = %d, want %d", s2, 66)
	}

	l1 := l * 100.0
	l2 := int(math.Round(l1))

	if l2 != 28 {
		t.Errorf("lum = %d, want %d", l2, 28)
	}
}

func TestHSLtoRGB(t *testing.T) {
	h := 193.0
	s := 0.67
	l := 0.28
	r, g, b := HSLtoRGB(h, s, l)
	if r != 23 {
		t.Errorf("red %d want %d", r, 23)
	}
	if g != 98 {
		t.Errorf("green %d want %d", g, 98)
	}
	if b != 119 {
		t.Errorf("blue %d want %d", b, 119)
	}
}

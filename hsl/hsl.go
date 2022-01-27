/* hsl
A library of routines to help to convert RGB to and from HSL
*/
package hsl

// RGBtoHSL : Convert RGB to HSL
// rgb values are 0-255
// h is between 0.0 - 360.0
// s & l values are 0.0 - 1.0
func RGBtoHSL(r, g, b uint8) (h, s, l float64) {

	var rgb [3]float64
	rgb[0] = float64(r) / 255.0
	rgb[1] = float64(g) / 255.0
	rgb[2] = float64(b) / 255.0

	max := 0.0
	for _, v := range rgb {
		if v > max {
			max = v
		}
	}

	min := 2.0
	for _, v := range rgb {
		if v < min {
			min = v
		}
	}

	l = (max + min) / 2.0

	if min == max {
		h = 0
		s = 0
	} else {
		d := max - min

		if l <= 0.5 {
			s = d / (max + min)
		} else {
			s = d / (2.0 - d)
		}

		if rgb[0] == max {
			h = (rgb[1] - rgb[2]) / d
		}
		if rgb[1] == max {
			h = 2.0 + (rgb[2]-rgb[0])/d
		}
		if rgb[2] == max {
			h = 4.0 + (rgb[0]-rgb[1])/d
		}

		h *= 60
		if h < 0 {
			h += 360.0
		}

	}
	return
}

// HSLtoRGB : Convert HSL to RGB format
// h is between 0 - 360
// s and l are between 0.0 - 1.0
// RGB is between 0-255
func HSLtoRGB(h, s, l float64) (r, g, b uint8) {
	var ttr, ttg, ttb float64

	if s == 0 {
		r = uint8(l * 255.0)
		g = uint8(l * 255.0)
		b = uint8(l * 255.0)
		return
	}

	var t1 float64
	if l < 0.5 {
		t1 = l * (1.0 + s)
	} else {
		t1 = l + s - l*s
	}

	t2 := 2*l - t1

	h /= 360.0

	var tr, tb, tg float64
	tr = h + 0.3333333333333333
	tg = h
	tb = h - 0.3333333333333333

	if tr < 0 {
		tr += 1
	}
	if tg < 0 {
		tg += 1
	}
	if tb < 0 {
		tb += 1
	}

	if tr > 1 {
		tr -= 1
	}
	if tg > 1 {
		tg -= 1
	}
	if tb > 1 {
		tb -= 1
	}

	if 6*tr < 1 {
		ttr = t2 + (t1-t2)*6*tr
	} else if 2*tr < 1 {
		ttr = t1
	} else if 3*tr < 2 {
		ttr = t2 + (t1-t2)*(0.66666-tr)*6
	} else {
		ttr = t2
	}

	if 6*tg < 1 {
		ttg = t2 + (t1-t2)*6*tg
	} else if 2*tg < 1 {
		ttg = t1
	} else if 3*tg < 2 {
		ttg = t2 + (t1-t2)*(0.66666-tg)*6
	} else {
		ttg = t2
	}

	if 6*tb < 1 {
		ttb = t2 + (t1-t2)*6*tb
	} else if 2*tb < 1 {
		ttb = t1
	} else if 3*tb < 2 {
		ttb = t2 + (t1-t2)*(0.66666-tb)*6
	} else {
		ttb = t2
	}

	r = uint8(ttr * 255.0)
	g = uint8(ttg * 255.0)
	b = uint8(ttb * 255.0)

	return
}

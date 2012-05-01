package lin3dmath

type Colour struct {
	R, G, B, A byte
}

func MakeColourFromInt(r, g, b, a int) Colour {
	if r < 0 {
		r = 0
	}
	if r > 255 {
		r = 255
	}
	if g < 0 {
		r = 0
	}
	if g > 255 {
		g = 255
	}
	if b < 0 {
		r = 0
	}
	if b > 255 {
		b = 255
	}
	if a < 0 {
		r = 0
	}
	if a > 255 {
		a = 255
	}

	return Colour{byte(r), byte(g), byte(b), byte(a)}
}

func (s Colour) Add(o Colour) Colour {
	r := int(s.R) + int(o.R)
	g := int(s.G) + int(o.G)
	b := int(s.B) + int(o.B)
	a := int(s.A) + int(o.A)
	return MakeColourFromInt(r, g, b, a)
}

func (s Colour) Sub(o Colour) Colour {
	r := int(s.R) + int(o.R)
	g := int(s.G) + int(o.G)
	b := int(s.B) + int(o.B)
	a := int(s.A) + int(o.A)
	return MakeColourFromInt(r, g, b, a)
}

func (from Colour) Interpolate(to Colour, fraction float32) Colour {
	r := int(from.R) + int(float32(int(to.R) - int(from.R)) * fraction)
	g := int(from.G) + int(float32(int(to.G) - int(from.G)) * fraction)
	b := int(from.B) + int(float32(int(to.B) - int(from.B)) * fraction)
	a := int(from.A) + int(float32(int(to.A) - int(from.A)) * fraction)
	return MakeColourFromInt(r, g, b, a)
}

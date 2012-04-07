package lin3dmath

import "math"

type Vector3f struct {
	X, Y, Z float32
}

func (s Vector3f) Neg() Vector3f {
	return Vector3f{-s.X, -s.Y, -s.Z}
}

func (s Vector3f) Normalize() Vector3f {
	dist2 := s.X*s.X + s.Y*s.Y + s.Z*s.Z
	dist := float32(math.Sqrt(float64(dist2)))

	return Vector3f{s.X / dist,
		s.Y / dist,
		s.Z / dist}
}

func (s *Vector3f) NormalizeIP() {
	dist2 := s.X*s.X + s.Y*s.Y + s.Z*s.Z
	dist := float32(math.Sqrt(float64(dist2)))

	s.X /= dist
	s.Y /= dist
	s.Z /= dist
}

func (s Vector3f) Len2() float32 {
	return s.X*s.X + s.Y*s.Y + s.Z*s.Z
}

func (s Vector3f) To2F() Vector2f {
	return Vector2f{float32(s.X),
		float32(s.Y)}
}

func (s Vector3f) To3I() Vector3i {
	return Vector3i{int(s.X),
		int(s.Y),
		int(s.Z)}
}

func (s Vector3f) Add(o Vector3f) Vector3f {
	return Vector3f{s.X + o.X, s.Y + o.Y, s.Z + o.Z}
}

func (s *Vector3f) AddIP(o Vector3f) {
	s.X += o.X; s.Y += o.Y; s.Z += o.Z
}

func (s Vector3f) Sub(o Vector3f) Vector3f {
	return Vector3f{s.X - o.X, s.Y - o.Y, s.Z - o.Z}
}

func (s Vector3f) Div(d float32) Vector3f {
	return Vector3f{s.X / d, s.Y / d, s.Z / d}
}

func (s Vector3f) Mul(d float32) Vector3f {
	return Vector3f{s.X * d, s.Y * d, s.Z * d}
}

func (s *Vector3f) MulIP(d float32) {
	s.X *= d
	s.Y *= d
	s.Z *= d
}

func (s Vector3f) Cross(o Vector3f) Vector3f {
	return Vector3f{s.Y*o.Z - o.Y*s.Z, s.Z*o.X - o.Z*s.X, s.X*o.Y - o.X*s.Y}
}

func (s Vector3f) Dot(o Vector3f) float32 {
	return s.X*o.X + s.Y*o.Y + s.Z*o.Z
}

func (s Vector3f) Len() float32 {
	return float32(math.Sqrt(float64(s.X*s.X + s.Y*s.Y + s.Z*s.Z)))
}

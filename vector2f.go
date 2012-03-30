package lin3dmath

import "math"

type Vector2f struct {
	X float32
	Y float32
}

func (s *Vector2f) NormalizeIP() {
	dist2 := s.X*s.X + s.Y*s.Y

	if dist2 != 0 {
		dist := float32(math.Sqrt(float64(dist2)))
		s.X /= dist
		s.Y /= dist
	}
}

func (s *Vector2f) Normalize() (r Vector2f) {
	r = *s
	r.NormalizeIP()
	return
}

func (s *Vector2f) AbsIP() {
	s.X = float32(math.Abs(float64(s.X)))
	s.Y = float32(math.Abs(float64(s.Y)))
}

func (s *Vector2f) Abs() (r Vector2f) {
	r = *s
	r.AbsIP()
	return
}

func (s *Vector2f) Len2() float32 {
	return s.X*s.X + s.Y*s.Y
}

func (s *Vector2f) To3F() Vector3f {
	return Vector3f{float32(s.X),
		float32(s.Y),
		0}
}

func (s *Vector2f) ToI() Vector2i {
	return Vector2i{int(s.X),
		int(s.Y)}
}
func (s *Vector2f) To2IT() (int, int) {
	return int(s.X), int(s.Y)
}

func (s *Vector2f) Add(o Vector2f) Vector2f {
	return Vector2f{s.X + o.X, s.Y + o.Y}
}

func (s *Vector2f) AddIP(o Vector2f) {
	s.X += o.X
	s.Y += o.Y
}

func (s *Vector2f) Sub(o Vector2f) Vector2f {
	return Vector2f{s.X - o.X, s.Y - o.Y}
}

func (s *Vector2f) SubIP(o Vector2f) {
	s.X -= o.X
	s.Y -= o.Y
}

func (s *Vector2f) Dot(o Vector2f) float32 {
	return s.X*o.X + s.Y*o.Y
}

func (s *Vector2f) Div(d float32) Vector2f {
	return Vector2f{s.X / d, s.Y / d}
}

func (s *Vector2f) DivIP(d float32) {
	s.X /= d
	s.Y /= d
}

func (s *Vector2f) Mul(d float32) Vector2f {
	return Vector2f{s.X * d, s.Y * d}
}

func (s *Vector2f) ZeroIP() {
	s.X = 0
	s.Y = 0
}

func (s *Vector2f) MulIP(d float32) {
	s.X *= d
	s.Y *= d
}

func (s *Vector2f) GetAngle() Angle {
	return Angle(float32(math.Atan2(float64(s.Y), float64(s.X))))
}

func (s *Vector2f) GetAngle64() float64 {
	return math.Atan2(float64(s.Y), float64(s.X))
}

func (s *Vector2f) Len() float32 {
	return float32(math.Sqrt(float64(s.X*s.X + s.Y*s.Y)))
}

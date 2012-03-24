package lin3dmath

type Vector2i struct {
	X,Y int
}

func (s *Vector2i) Add(o Vector2i) (Vector2i) {
	return Vector2i{s.X+o.X, s.Y+o.Y}
}

func (s *Vector2i) Sub(o Vector2i) (Vector2i) {
	return Vector2i{s.X-o.X, s.Y-o.Y}
}

func (s *Vector2i) Dot(o Vector2i) (int) {
	return s.X*o.X+s.Y*o.Y
}

func (s *Vector2i) Div(d int) (Vector2i) {
	return Vector2i{s.X/d, s.Y/d}
}

func (s *Vector2i) Mul(d int) (Vector2i) {
	return Vector2i{s.X*d, s.Y*d}
}

func (s *Vector2i) ZeroIP() {
	s.X = 0
	s.Y = 0
}

func (s *Vector2i) MulIP(d int) {
	s.X*=d
	s.Y*=d
}


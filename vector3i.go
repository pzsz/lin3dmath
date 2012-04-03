package lin3dmath

type Vector3i struct {
	X, Y, Z int
}

func (s Vector3i) Add(o Vector3i) Vector3i {
	return Vector3i{s.X + o.X, s.Y + o.Y, s.Z + o.Z}
}

func (s *Vector3i) AddIP(o Vector3i) {
	s.X += o.X; s.Y += o.Y; s.Z += o.Z;
}

func (s Vector3i) MulI(o int) Vector3i {
	return Vector3i{s.X*o, s.Y*o, s.Z*o}
}

func (s Vector3i) Mul3I(o Vector3i) Vector3i {
	return Vector3i{s.X*o.X, s.Y*o.Y, s.Z*o.Z}
}

func (s Vector3i) ToF() Vector3f {
	return Vector3f{float32(s.X),
		float32(s.Y),
		float32(s.Z)}
}

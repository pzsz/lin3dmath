package lin3dmath

type Boxi struct {
	Start  Vector3i
	End    Vector3i
} 

func (s Boxi) GrowBy(val int) Boxi {
	s.Start.X -= val
	s.Start.Y -= val
	s.Start.Z -= val

	s.End.X += val
	s.End.Y += val
	s.End.Z += val
	return s
}

func (s *Boxi) Intersection(b Boxi) (exist bool, r Boxi) {
	exist = true
	if s.Start.X >= b.Start.X {
		r.Start.X = s.Start.X
	} else {
		r.Start.X = b.Start.X
	}
	if s.End.X <= b.End.X {
		r.End.X = s.End.X
	} else {
		r.End.X = b.End.X
	}
	if r.End.X < r.Start.X {
		exist = false
		return
	}

	if s.Start.Y >= b.Start.Y {
		r.Start.Y = s.Start.Y
	} else {
		r.Start.Y = b.Start.Y
	}
	if s.End.Y <= b.End.Y {
		r.End.Y = s.End.Y
	} else {
		r.End.Y = b.End.Y
	}
	if r.End.Y < r.Start.Y {
		exist = false
		return
	}


	if s.Start.Z >= b.Start.Z {
		r.Start.Z = s.Start.Z
	} else {
		r.Start.Z = b.Start.Z
	}
	if s.End.Z <= b.End.Z {
		r.End.Z = s.End.Z
	} else {
		r.End.Z = b.End.Z
	}
	if r.End.Z < r.Start.Z {
		exist = false
		return
	}
	return
}
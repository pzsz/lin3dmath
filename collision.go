package lin3dmath

func TestCollisionCircleVsBox(c1 Circle, b2 Rect) (isCol bool, cont *Vector2f) {
	isCol = false
	Radius := c1.Radius
	Radius2 := Radius *Radius
	Pos := c1.Pos

	// Place test point in center of cord system
	blockDownX := b2.Pos.X
	blockDownY := b2.Pos.Y
	blockTopX  := b2.Pos.X + b2.Width
	blockTopY  := b2.Pos.Y + b2.Height

	blockMinX := blockDownX - c1.Pos.X;
	blockMinY := blockDownY - c1.Pos.Y;
	blockMaxX := blockTopX - c1.Pos.X;
	blockMaxY := blockTopY - c1.Pos.Y;

	// Inside
	if ((blockDownX <= Pos.X && blockTopX >= Pos.X) && (blockDownY <= Pos.Y && blockTopY >= Pos.Y)) {
		return true,&Vector2f{Pos.X, Pos.Y};
	}

	if (blockMaxX < 0) {
		if (blockMaxY < 0) {
			if ((blockMaxX * blockMaxX + blockMaxY * blockMaxY) < Radius2) {
				return true,&Vector2f{blockTopX, blockTopY};
			}
		} else if (blockMinY > 0) {
			if ((blockMaxX * blockMaxX + blockMinY * blockMinY) < Radius2) {
				return true,&Vector2f{blockTopX, blockDownY};
			}
		} else {
			if (-blockMaxX < Radius) {
				return true,&Vector2f{blockTopX, Pos.Y};
			}
		}
	} else if (blockMinX > 0) {
		if (blockMaxY < 0) {
			if ((blockMinX * blockMinX + blockMaxY * blockMaxY) < Radius2) {
				return true,&Vector2f{blockDownX, blockTopY};
			}
		} else if (blockMinY > 0) {
			if ((blockMinX * blockMinX + blockMinY * blockMinY) < Radius2) {
				return true,&Vector2f{blockDownX, blockDownY};
			}
		} else {
			if (blockMinX <= Radius) {
				return true,&Vector2f{blockDownX, Pos.Y};
			}
		}
	} else {
		if (blockMaxY < 0) {
			if (-blockMaxY <= Radius) {
				return true,&Vector2f{Pos.X, blockTopY};
			}
		} else if (blockMinY > 0) {
			if (blockMinY <= Radius){
				return true,&Vector2f{Pos.X, blockDownY};
			}
		}
	}

	return false, nil;
}

func TestCollisionCircleVsCircle(c1 Circle, c2 Circle) (isCol bool, cont *Vector2f) {
	dist := c1.Pos.Sub(c2.Pos)
	dlen := dist.Len()

	diff := dlen - c1.Radius - c2.Radius

	if diff > 0 {				
		return false,nil
	}
	
	c2rat := c2.Radius / (c1.Radius + c2.Radius)

	half_dist := dist.Mul(c2rat)
	ret := c2.Pos.Add(half_dist)
	return true, &ret
}

func TestPointSideOnLine(l *Line, p Vector2f) int {
	s := (l.To.X - l.From.X) * (p.Y - l.From.Y) - (l.To.Y - l.From.Y) * (p.X - l.From.X)
	if s > 0 {return 1}
	if s < 0 {return -1}
	return 0
}

func TestCollisionLineVsBox(l Line, b2 Rect) bool  {
	s1 := TestPointSideOnLine(&l, b2.GetP1())
	s2 := TestPointSideOnLine(&l, b2.GetP2())
	s3 := TestPointSideOnLine(&l, b2.GetP3())
	s4 := TestPointSideOnLine(&l, b2.GetP4())

	s := s1 + s2 + s3 + s4
	
	return s == 4 || s != -4
}
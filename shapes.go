package lin3dmath

type Circle struct {
	Pos    Vector2f
	Radius float32
}

func MakeCircle(x, y, radius float32) Circle {
	return Circle{Vector2f{x, y}, radius}
}

type Rect struct {
	Pos    Vector2f
	Width  float32
	Height float32
}

func (self *Rect) SetSize(v Vector2f) {
	self.Width = v.X
	self.Height = v.Y
}

func (self *Rect) GetStartX() float32 {
	return self.Pos.X
}

func (self *Rect) GetStartY() float32 {
	return self.Pos.Y
}

func (self *Rect) GetEndX() float32 {
	return self.Pos.X + self.Width
}

func (self *Rect) GetEndY() float32 {
	return self.Pos.Y + self.Height
}

func (self *Rect) GetP1() Vector2f {
	return self.Pos
}

func (self *Rect) GetP2() Vector2f {
	return Vector2f{self.Pos.X + self.Width, self.Pos.Y}
}

func (self *Rect) GetP3() Vector2f {
	return Vector2f{self.Pos.X + self.Width, self.Pos.Y + self.Height}
}

func (self *Rect) GetP4() Vector2f {
	return Vector2f{self.Pos.X, self.Pos.Y + self.Height}
}

func (self *Rect) IsInside(x, y float32) bool {
	return x >= self.Pos.X && x <= (self.Pos.X+self.Width) &&
		y >= self.Pos.Y && y <= (self.Pos.Y+self.Height)
}

type Rect2i struct {
	Pos    Vector2i
	Width  int
	Height int
}

func (self *Rect2i) SetSize(v Vector2i) {
	self.Width = v.X
	self.Height = v.Y
}

func (self *Rect2i) GetStartX() int {
	return self.Pos.X
}

func (self *Rect2i) GetStartY() int {
	return self.Pos.Y
}

func (self *Rect2i) GetEndX() int {
	return self.Pos.X + self.Width
}

func (self *Rect2i) GetEndY() int {
	return self.Pos.Y + self.Height
}

func (self *Rect2i) GetP1() Vector2i {
	return self.Pos
}

func (self *Rect2i) GetP2() Vector2i {
	return Vector2i{self.Pos.X + self.Width, self.Pos.Y}
}

func (self *Rect2i) GetP3() Vector2i {
	return Vector2i{self.Pos.X + self.Width, self.Pos.Y + self.Height}
}

func (self *Rect2i) GetP4() Vector2i {
	return Vector2i{self.Pos.X, self.Pos.Y + self.Height}
}

func (self *Rect2i) IsInside(x, y int) bool {
	return x >= self.Pos.X && x <= (self.Pos.X+self.Width) &&
		y >= self.Pos.Y && y <= (self.Pos.Y+self.Height)
}

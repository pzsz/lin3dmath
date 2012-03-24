package lin3dmath 

import "math"

type Degree float32
type Angle float32

func NewAngleFromVector(v Vector2f) Angle {
	return Angle(math.Atan2(float64(v.Y),float64(v.X)))
}

func (self Degree) ToAngle() Angle {
	return Angle((self * 2*math.Pi) / 360.0)
}

func (self Angle) ToDegree() Degree {
	return Degree((self * 360.0) / (2*math.Pi))
}

func (self Angle) X() float32 {
	return float32(math.Cos(float64(self)))
}

func (self Angle) Y() float32 {
	return float32(math.Sin(float64(self)))
}
func (self *Degree) Normalize() {
	for ;*self < -180.0;*self += 360 {}
	for ;*self >  180.0;*self -= 360 {}
}

func (self Degree) X() float32 {
	return float32(math.Cos(float64(self.ToAngle())))
}

func (self Degree) Y() float32 {
	return float32(math.Sin(float64(self.ToAngle())))
}
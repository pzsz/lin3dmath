package lin3dmath

import (
	"math"
)

type Matrix4 [16]float32

func Matrix4Zero() *Matrix4 {
	return &Matrix4{}
}

func MatrixOne() *Matrix4 {
	return &Matrix4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1}
}

func MatrixTranslate(x, y, z float32) *Matrix4 {
	return &Matrix4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		x, y, z, 1}
}

func MatrixScale(x, y, z float32) *Matrix4 {
	return &Matrix4{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1}
}

func MatrixRotate(angle Angle, x, y, z float32) *Matrix4 {
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))

	return &Matrix4{
		x*x*(1-c) + c, x*y*(1-c) + z*s, x*z*(1-c) - y*s, 0,
		y*x*(1-c) - z*s, y*y*(1-c) + c, y*z*(1-c) + x*s, 0,
		x*z*(1-c) + y*s, y*z*(1-c) - x*s, z*z*(1-c) + c, 0,
		0, 0, 0, 1}

}

func (a *Matrix4) Mul(b *Matrix4) (ret Matrix4) {
	for c := 0; c < 4; c++ {
		for r := 0; r < 4; r++ {
			ret[r*4+c] = b[r*4+0]*a[0+c] +
				b[r*4+1]*a[4+c] +
				b[r*4+2]*a[8+c] +
				b[r*4+3]*a[12+c]
		}
	}
	return
}

func (a *Matrix4) ToSlice() []float32 {
	return a[0:16]
}

func (a *Matrix4) ToSlice64() []float64 {
	var ret [16]float64
	for x := 0; x < 16; x++ {
		ret[x] = float64(a[x])
	}
	return ret[0:16]
}

func (a *Matrix4) ToArray32() *[16]float32 {
	return (*[16]float32)(a)
}

func (a *Matrix4) ToArray64() *[16]float64 {
	var ret [16]float64
	for x := 0; x < 16; x++ {
		ret[x] = float64(a[x])
	}
	return &ret
}

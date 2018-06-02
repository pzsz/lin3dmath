package lin3dmath

type Quaternion struct {
	X, Y, Z, W float32
}

func QuaternionFromAngle(axis Vector3f, angle Angle) Quaternion {
	d := axis.Len()
	if d == 0 {
		panic("Axis of length zero!")
	}
	s := Sin32(float32(angle)*0.5) / d

	return Quaternion{axis.X * s, axis.Y * s, axis.Z * s, Cos32(float32(angle) * 0.5)}
}

func (q *Quaternion) Scale(val float32) *Quaternion {
	return &Quaternion{
		q.X * val,
		q.Y * val,
		q.Z * val,
		q.W * val,
	}
}

func (q *Quaternion) Normalize() *Quaternion {
	return q.Scale(1.0 / q.Length())
}

func (q *Quaternion) Invert() *Quaternion {
	return (&Quaternion{
		-q.X,
		-q.Y,
		-q.Z,
		q.W,
	}).Normalize()
}

func (q *Quaternion) Length2() float32 {
	return q.X*q.X + q.Y*q.Y + q.Z*q.Z + q.W*q.W
}

func (q *Quaternion) Length() float32 {
	return Sqrt32(q.X*q.X + q.Y*q.Y + q.Z*q.Z + q.W*q.W)
}

func (q *Quaternion) Mul(o *Quaternion) *Quaternion {
	return &Quaternion{
		q.W*o.X + q.X*o.W + q.Y*o.Z - q.Z*o.Y,
		q.W*o.Y + q.Y*o.W + q.Z*o.X - q.X*o.Z,
		q.W*o.Z + q.Z*o.W + q.X*o.Y - q.Y*o.X,
		q.W*o.W - q.X*o.X - q.Y*o.Y - q.Z*o.Z}
}

func (a *Quaternion) MulVector3f(b Vector3f) *Quaternion {
	return &Quaternion{
		a.W*b.X + a.Y*b.Z - a.Z*b.Y,
		a.W*b.Y - a.X*b.Z + a.Z*b.X,
		a.W*b.Z + a.X*b.Y - a.Y*b.X,
		-a.X*b.X - a.Y*b.Y - a.Z*b.Z}
}

func (q *Quaternion) TransformVector3f(v Vector3f) Vector3f {
	res := q.MulVector3f(v)
	res = res.Mul(q.Invert())
	return Vector3f{
		res.X, res.Y, res.Z}
}

func (q *Quaternion) Mul3f(v Vector3f) Vector3f {
	qvec := Vector3f{q.X, q.Y, q.Z}
	uv := qvec.Cross(v)
	uuv := qvec.Cross(uv)
	uv.MulIP(2.0 * q.W)
	uuv.MulIP(2.0)

	return Vector3f{
		v.X + uv.X + uuv.X,
		v.Y + uv.Y + uuv.Y,
		v.Z + uv.Z + uuv.Z}
}

func (q *Quaternion) GetAngle() Angle {
	return Angle(2.0 * Acos32(q.W))
}

func (q *Quaternion) ToMatrix4() *Matrix4 {
	d := q.Length2()
	if d == 0 {
		panic("Quaternion zero!")
	}
	s := 2.0 / d
	xs := q.X * s
	ys := q.Y * s
	zs := q.Z * s
	wx := q.W * xs
	wy := q.W * ys
	wz := q.W * zs
	xx := q.X * xs
	xy := q.X * ys
	xz := q.X * zs
	yy := q.Y * ys
	yz := q.Y * zs
	zz := q.Z * zs

	return &Matrix4{
		1.0 - (yy + zz), xy - wz, xz + wy, 0,
		xy + wz, 1.0 - (xx + zz), yz - wx, 0,
		xz - wy, yz + wx, 1.0 - (xx + yy), 0,
		0, 0, 0, 1}
}

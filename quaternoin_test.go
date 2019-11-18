package lin3dmath_test

import (
	"testing"

	. "github.com/finalist736/lin3dmath"
)

func TestQVMul(t *testing.T) {
	v := Vector3f{0.0, 0.0, 1.0}
	q := QuaternionFromAngle(Vector3f{1.0, 0.0, 0.0}, Degree(90).ToAngle())
	v2 := q.TransformVector3f(v)
	t.Logf("1 %+v\n", v2)
	if v2.X != 0.0 || v2.Y != -1.0 || v2.Z != 0.0 {
		t.Errorf("Transform error: Vector: %+v, \n\tQuaternoin: %+v, \n\tresult: %+v", v, q, v2)
	} //0.707107
	q = QuaternionFromAngle(Vector3f{1.0, 0.0, 0.0}, Degree(45).ToAngle())
	v2 = q.TransformVector3f(v)
	t.Logf("2 %+v\n", v2)
	t.Logf("3 %+v, %+v\n", RoundUnit(v2.Y, 0.0005), RoundUnit(v2.Z, 0.0005))
	if v2.X != 0.0 || RoundUnit(v2.Y, 0.0005) != -0.707 || RoundUnit(v2.Z, 0.0005) != 0.707 {
		t.Errorf("Transform error: Vector: %+v, \n\tQuaternoin: %+v, \n\tresult: %+v", v, q, v2)
	} //0.707107
}

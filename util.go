package lin3dmath

import "math"

type Direction int

const (
	UP    = Direction(1)
	NORTH = UP
	DOWN  = Direction(2)
	SOUTH = DOWN
	LEFT  = Direction(4)
	WEST  = LEFT
	RIGHT = Direction(8)
	EAST  = RIGHT

	UP_RIGHT   = UP | RIGHT
	DOWN_RIGHT = DOWN | RIGHT

	UP_LEFT   = UP | LEFT
	DOWN_LEFT = DOWN | LEFT
)

func (s Direction) ToVector2f() Vector2f {
	switch s {
	case UP:
		return Vector2f{0, 1}
	case DOWN:
		return Vector2f{0, -1}
	case LEFT:
		return Vector2f{-1, 0}
	case RIGHT:
		return Vector2f{1, 0}
	}
	panic("Unknown direction")
}

func (s Direction) ToVector2i() Vector2i {
	switch s {
	case UP:
		return Vector2i{0, 1}
	case DOWN:
		return Vector2i{0, -1}
	case LEFT:
		return Vector2i{-1, 0}
	case RIGHT:
		return Vector2i{1, 0}
	}
	panic("Unknown direction")
}

func (s Direction) IsVertical() bool {
	return s == UP || s == DOWN
}

func (s Direction) IsHorizontal() bool {
	return s == LEFT || s == RIGHT
}

func AlignDivUpI(v, align int) int {
	return (v + align - 1) / align
}

func AlignUpF(v, align int) int {
	return (v + align - 1) / align
}

func AlignDivDown2FToI(v Vector2f, align_x, align_y int) (x, y int) {
	x = int(v.X / float32(align_x))
	y = int(v.Y / float32(align_y))
	return
}

func ValidateCoordRangeI(x, y, w, h int) bool {
	return x >= 0 && y >= 0 && x < w && y < h
}

func RoundF(f float32) float32 {
	fi := float32(int(f))
	dif := f - fi
	return fi + dif*2
}

func DistToNextCell(x, y float32, dir Direction) float32 {
	ix := float32(int(x))
	iy := float32(int(y))

	switch dir {
	case UP:
		return iy + 1 - y
	case RIGHT:
		return ix + 1 - x
	case DOWN:
		if y == iy {
			return -1
		}
		return iy - y
	case LEFT:
		if x == ix {
			return -1
		}
		return ix - x
	}
	panic("Unknown dir")
}

func AlignToCell(val float32, dir Direction) int {
	iv := int(val)
	fiv := float32(iv)

	if (dir == UP || dir == RIGHT) || fiv != val {
		return iv
	}

	return iv - 1
}

func Abs32(v float32) float32 {
	return float32(math.Abs(float64(v)))
}

func FastFloor64(v float64) float64 {
	i := float64(int(v))
	if v < i {
		return i-1
	}
	return i
}

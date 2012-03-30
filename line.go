package lin3dmath

import (
//	"fmt"
)

type Line struct {
	From Vector2f
	To   Vector2f
}

func (self *Line) GetDirVector() Vector2f {
	return self.To.Sub(self.From)
}

func (self *Line) GetDistancesVector() Vector2f {
	dir := self.GetDirVector()
	dir.AbsIP()
	return dir
}

type RasterizeCallback func(x, y int) bool

func (self *Line) Rasterize(callback RasterizeCallback) {
	dir := self.GetDirVector()

	if dir.X == 0 && dir.Y == 0 {
		return
	}
	dir.NormalizeIP()

	// Vertical case
	if dir.X == 0 {
		x, sy, ey, step := int(self.To.X), int(self.From.Y), int(self.To.Y), int(dir.Y)
		for y := sy; y != ey; y += step {
			if !callback(x, y) {
				return
			}
		}
		return
	}

	// Horizontal case
	if dir.Y == 0 {
		y, sx, ex, step := int(self.To.Y), int(self.From.X), int(self.To.X), int(dir.X)
		for x := sx; x != ex; x += step {
			if !callback(x, y) {
				return
			}
		}
		return
	}

	xRatio := dir.Y / dir.X
	yRatio := dir.X / dir.Y

	// Get dir
	var xDir, yDir Direction
	if dir.X >= 0 {
		xDir = RIGHT
	} else {
		xDir = LEFT
	}
	if dir.Y >= 0 {
		yDir = UP
	} else {
		yDir = DOWN
	}

	//fmt.Printf("Start %v %v\n", self.From, self.To)

	xPos, yPos := self.From.X, self.From.Y

	distances := self.GetDistancesVector()

	for {
		if distances.X < 0 && distances.Y < 0 {
			//fmt.Printf("Win\n")
			return
		}

		rasterX := calcRasterCoord(xPos, xDir)
		rasterY := calcRasterCoord(yPos, yDir)

		//fmt.Printf("Pass %v %v (%v %v)\n", xPos, yPos, rasterX, rasterY)

		if !callback(rasterX, rasterY) {
			//fmt.Printf("Fail\n")
			return
		}

		xToAlign := DistToNextCell(xPos, yPos, xDir)
		yAfterAlign := xToAlign * xRatio

		if AlignToCell(yPos+yAfterAlign, yDir) == AlignToCell(yPos, yDir) {
			//fmt.Printf("Align to X\n")
			newXPos := RoundF(xToAlign + xPos)

			distances.X -= Abs32(newXPos - xPos)
			distances.Y -= Abs32(yAfterAlign)

			yPos += yAfterAlign
			xPos = newXPos
		} else {
			//fmt.Printf("Align to Y\n")
			yToAlign := DistToNextCell(xPos, yPos, yDir)
			xAfterAlign := yToAlign * yRatio
			newYPos := RoundF(yToAlign + yPos)

			distances.Y -= Abs32(newYPos - yPos)
			distances.X -= Abs32(xAfterAlign)

			xPos += xAfterAlign
			yPos = newYPos
		}
	}
}

/* Calculate raster position we just went through */
func calcRasterCoord(pos float32, dir Direction) int {
	iv := int(pos)
	fiv := float32(iv)

	if fiv == pos {
		if dir == RIGHT || dir == UP {
			return iv - 1
		}
		return iv
	}

	return iv
}

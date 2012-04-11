package lin3dmath

import (
	"math/rand"
	"math"
)

type PerlinNoise3D struct {
	seed   int64
	permut [256]int
	g3d    [256][3]float64 // Randomly generated 2D unit vectors.
}

func NewPerlinNoise3D(seed int64) *PerlinNoise3D {
	gen := &PerlinNoise3D{
		seed: seed,
	}

	// The source's seed is reset to seed for each precomputed set so that any
	// code reordering in this implementation does not alter the noise values
	// produced for a given seed.
	source := rand.NewSource(0)
	rnd := rand.New(source)

	// Initialize gen.permut.
	source.Seed(seed)
	perm := rnd.Perm(len(gen.permut))
	for i := range perm {
		gen.permut[i] = perm[i]
	}

	// Initialize gen.g2d.
	source.Seed(seed)
	for i := range perm {
		length := 0.0
		for j:=0; j<3; j++ {
			gen.g3d[i][j] = 2*rnd.Float64() - 1
			length += gen.g3d[i][j]*gen.g3d[i][j]
		}
		// Normalize
		length = math.Sqrt(length)
		for j:=0; j<3; j++ {
			gen.g3d[i][j] /= length
		}
	}

	return gen
}

func (gen *PerlinNoise3D) grad3d(x, y, z int) *[3]float64 {
	gradIndex := gen.permut[x&0xff] + gen.permut[y&0xff] + gen.permut[z&0xff]
	return &gen.g3d[gradIndex&0xff]
}

// At2d returns the noise value at a given 2D point.
func (gen *PerlinNoise3D) At(x, y, z float64) float64 {
	x0 := FastFloor64(x)
	y0 := FastFloor64(y)
	z0 := FastFloor64(z)
	x1 := x0 + 1
	y1 := y0 + 1
	z1 := z0 + 1

	// Label corners S=x0,y0, T=x1,y0, U=x0,y1, V=x1,y1.
	gradSL := gen.grad3d(int(x0), int(y0), int(z0))
	gradTL := gen.grad3d(int(x1), int(y0), int(z0))
	gradUL := gen.grad3d(int(x0), int(y1), int(z0))
	gradVL := gen.grad3d(int(x1), int(y1), int(z0))

	// Label corners S=x0,y0, T=x1,y0, U=x0,y1, V=x1,y1.
	gradSU := gen.grad3d(int(x0), int(y0), int(z1))
	gradTU := gen.grad3d(int(x1), int(y0), int(z1))
	gradUU := gen.grad3d(int(x0), int(y1), int(z1))
	gradVU := gen.grad3d(int(x1), int(y1), int(z1))

	// dotX := gradX · ((x,y) - (xX,yX))
	dotSL := gradSL[0]*(x-x0) + gradSL[1]*(y-y0) + gradSL[2]*(z-z0) 
	dotTL := gradTL[0]*(x-x1) + gradTL[1]*(y-y0) + gradTL[2]*(z-z0) 
	dotUL := gradUL[0]*(x-x0) + gradUL[1]*(y-y1) + gradUL[2]*(z-z0) 
	dotVL := gradVL[0]*(x-x1) + gradVL[1]*(y-y1) + gradVL[2]*(z-z0) 
        
	dotSU := gradSU[0]*(x-x0) + gradSU[1]*(y-y0) + gradSU[2]*(z-z1) 
	dotTU := gradTU[0]*(x-x1) + gradTU[1]*(y-y0) + gradTU[2]*(z-z1) 
	dotUU := gradUU[0]*(x-x0) + gradUU[1]*(y-y1) + gradUU[2]*(z-z1) 
	dotVU := gradVU[0]*(x-x1) + gradVU[1]*(y-y1) + gradVU[2]*(z-z1) 

	// Bilinear interpolation of the weight between all four points, but using an
	// "ease" function.
	dx := x - x0
	dy := y - y0
	dz := z - z0

	sx := 3*dx*dx - 2*dx*dx*dx
	sy := 3*dy*dy - 2*dy*dy*dy 
	sz := 3*dz*dz - 2*dz*dz*dz

	al := dotSL + sx*(dotTL-dotSL)
	bl := dotUL + sx*(dotVL-dotUL)
	l := al + sy*(bl-al)
	
	au := dotSU + sx*(dotTU-dotSU)
	bu := dotUU + sx*(dotVU-dotUU)
	u := au + sy*(bu-au)

	return l + sz*(u-l)
}
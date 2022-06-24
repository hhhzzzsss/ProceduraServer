package region

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"sort"
	"unicode"

	"github.com/hhhzzzsss/procedura-generator/util"
)

const temp_path string = "TEMP_REGION_DUMP"
const output_path string = "../plugins/ProceduraPlugin/PROCEDURA_REGION_DUMP"

type Region struct {
	xdim, ydim, zdim int
	ids              []int
	palette          []string
}

func MakeRegion(xdim, ydim, zdim int) Region {
	return Region{
		xdim, ydim, zdim,
		make([]int, xdim*ydim*zdim),
		make([]string, 0),
	}
}

func (r *Region) AddPaletteBlock(block string) {
	r.palette = append(r.palette, block)
}

func (r *Region) PaletteSize() int {
	return len(r.palette)
}

func (r *Region) Set(x, y, z, id int) {
	if !r.IsInRange(x, y, z) {
		return
	}
	r.ids[y*r.zdim*r.xdim+z*r.xdim+x] = id
}

func (r *Region) Get(x, y, z int) int {
	if !r.IsInRange(x, y, z) {
		return 0
	}
	return r.ids[y*r.zdim*r.xdim+z*r.xdim+x]
}

func (r *Region) ForEach(idGenerator func(x, y, z int) int) {
	var bar util.ProgressBar
	bar.Initialize(r.ydim)
	for y := 0; y < r.ydim; y++ {
		for z := 0; z < r.zdim; z++ {
			for x := 0; x < r.xdim; x++ {
				r.Set(x, y, z, idGenerator(x, y, z))
			}
		}
		bar.Play(y + 1)
	}
	bar.Finish()
}

func (r *Region) ForEachNormalized(idGenerator func(x, y, z float64) int) {
	var bar util.ProgressBar
	bar.Initialize(r.ydim)
	minDim := r.xdim
	if r.ydim < minDim {
		minDim = r.ydim
	}
	if r.zdim < minDim {
		minDim = r.zdim
	}
	for y := 0; y < r.ydim; y++ {
		for z := 0; z < r.zdim; z++ {
			for x := 0; x < r.xdim; x++ {
				xNorm := 2.0*float64(x)/float64(r.xdim) - 1.0
				yNorm := 2.0*float64(y)/float64(r.ydim) - 1.0
				zNorm := 2.0*float64(z)/float64(r.zdim) - 1.0
				xNorm *= float64(r.xdim) / float64(minDim)
				yNorm *= float64(r.ydim) / float64(minDim)
				zNorm *= float64(r.zdim) / float64(minDim)
				r.Set(x, y, z, idGenerator(xNorm, yNorm, zNorm))
			}
		}
		bar.Play(y + 1)
	}
	bar.Finish()
}

func (r *Region) ForEachInSphere(cx, cy, cz, radius float64, f func(x, y, z int, rad2 float64)) {
	x1 := int(math.Floor(cx - radius))
	y1 := int(math.Floor(cy - radius))
	z1 := int(math.Floor(cz - radius))
	x2 := int(math.Floor(cx + radius))
	y2 := int(math.Floor(cy + radius))
	z2 := int(math.Floor(cz + radius))
	for by := y1; by <= y2; by++ {
		for bz := z1; bz <= z2; bz++ {
			for bx := x1; bx <= x2; bx++ {
				dx := float64(bx) + 0.5 - cx
				dy := float64(by) + 0.5 - cy
				dz := float64(bz) + 0.5 - cz
				rad2 := dx*dx + dy*dy + dz*dz
				if rad2 <= radius*radius {
					f(bx, by, bz, rad2)
				}
			}
		}
	}
}

type sphereCandidate struct {
	x, y, z    int
	rad        float64
	sortFactor float64
}

func (r *Region) ForEachInVolumePreservingSphere(cx, cy, cz, radius float64, bias util.Vec3d, f func(x, y, z int, rad2 float64)) {
	x1 := int(math.Floor(cx-radius)) - 1
	y1 := int(math.Floor(cy-radius)) - 1
	z1 := int(math.Floor(cz-radius)) - 1
	x2 := int(math.Floor(cx+radius)) + 1
	y2 := int(math.Floor(cy+radius)) + 1
	z2 := int(math.Floor(cz+radius)) + 1
	candidates := make([]sphereCandidate, 0)
	for by := y1; by <= y2; by++ {
		for bz := z1; bz <= z2; bz++ {
			for bx := x1; bx <= x2; bx++ {
				dx := float64(bx) + 0.5 - cx
				dy := float64(by) + 0.5 - cy
				dz := float64(bz) + 0.5 - cz
				dvec := util.MakeVec3d(dx, dy, dz)
				rad := dvec.Length()
				if rad <= radius+0.5 {
					candidates = append(candidates, sphereCandidate{bx, by, bz, rad, rad + 0.5*math.Abs(dvec.Dot(bias))})
				}
			}
		}
	}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].sortFactor < candidates[j].sortFactor
	})
	targetVolume := 4. / 3. * math.Pi * radius * radius * radius
	currentVolume := 0.
	for _, candidate := range candidates {
		f(candidate.x, candidate.y, candidate.z, candidate.rad*candidate.rad)
		currentVolume++
		if currentVolume >= targetVolume {
			break
		}
	}
}

func (r *Region) ForEachOnBorder(f func(x, y, z int)) {
	for y := 0; y < r.ydim; y++ {
		for x := 0; x < r.xdim; x++ {
			f(x, y, 0)
			f(x, y, r.zdim-1)
		}
	}
	for y := 0; y < r.ydim; y++ {
		for z := 1; z < r.zdim-1; z++ {
			f(0, y, z)
			f(r.xdim-1, y, z)
		}
	}
	for x := 1; x < r.xdim-1; x++ {
		for z := 1; z < r.zdim-1; z++ {
			f(x, 0, z)
			f(x, r.ydim-1, z)
		}
	}
}

func (r *Region) Hollow() {
	fmt.Println("Hollowing...")
	isSurface := MakeRegionCache[bool](r)
	for y := 1; y < r.ydim-1; y++ {
		for z := 1; z < r.zdim-1; z++ {
			for x := 1; x < r.xdim-1; x++ {
				if r.Get(x+1, y, z) == 0 ||
					r.Get(x-1, y, z) == 0 ||
					r.Get(x, y+1, z) == 0 ||
					r.Get(x, y-1, z) == 0 ||
					r.Get(x, y, z+1) == 0 ||
					r.Get(x, y, z-1) == 0 {
					isSurface.Set(x, y, z, true)
				}
			}
		}
	}
	for y := 1; y < r.ydim-1; y++ {
		for z := 1; z < r.zdim-1; z++ {
			for x := 1; x < r.xdim-1; x++ {
				if !isSurface.Get(x, y, z) {
					r.Set(x, y, z, 0)
				}
			}
		}
	}
}

func (r *Region) SelectiveHollow(id int) {
	fmt.Printf("Hollowing blocks with id %d...\n", id)
	isSurface := MakeRegionCache[bool](r)
	for y := 1; y < r.ydim-1; y++ {
		for z := 1; z < r.zdim-1; z++ {
			for x := 1; x < r.xdim-1; x++ {
				if r.Get(x, y, z) != id {
					continue
				}
				if r.Get(x+1, y, z) != id ||
					r.Get(x-1, y, z) != id ||
					r.Get(x, y+1, z) != id ||
					r.Get(x, y-1, z) != id ||
					r.Get(x, y, z+1) != id ||
					r.Get(x, y, z-1) != id {
					isSurface.Set(x, y, z, true)
				}
			}
		}
	}
	for y := 1; y < r.ydim-1; y++ {
		for z := 1; z < r.zdim-1; z++ {
			for x := 1; x < r.xdim-1; x++ {
				if r.Get(x, y, z) != id {
					continue
				}
				if !isSurface.Get(x, y, z) {
					r.Set(x, y, z, 0)
				}
			}
		}
	}
}

func (r *Region) CreateDump() {
	r.Validate()

	fmt.Println("Writing region file...")
	f, err := os.Create(temp_path)
	if err != nil {
		panic(err)
	}

	binary.Write(f, binary.BigEndian, uint32(r.xdim))
	binary.Write(f, binary.BigEndian, uint32(r.ydim))
	binary.Write(f, binary.BigEndian, uint32(r.zdim))

	binary.Write(f, binary.BigEndian, uint32(len(r.palette)))
	for _, paletteStr := range r.palette {
		for _, c := range paletteStr {
			if c > unicode.MaxASCII {
				panic("Palette entry was not ascii")
			}
		}
		binary.Write(f, binary.BigEndian, uint32(len(paletteStr)))
		binary.Write(f, binary.BigEndian, []byte(paletteStr))
	}
	dataBuffer := make([]byte, r.xdim*r.ydim*r.zdim*4)
	for y := 0; y < r.ydim; y++ {
		for z := 0; z < r.zdim; z++ {
			for x := 0; x < r.xdim; x++ {
				bufferIdx := y*r.zdim*r.xdim*4 + z*r.xdim*4 + x*4
				binary.BigEndian.PutUint32(dataBuffer[bufferIdx:], uint32(r.Get(x, y, z)))
			}
		}
	}
	binary.Write(f, binary.BigEndian, dataBuffer)

	f.Close()
	os.Rename(temp_path, output_path)
	fmt.Println("Finished creating region file")
}

// Panics if region has invalid state
func (r *Region) Validate() {
	fmt.Println("Validating region...")
	for _, paletteStr := range r.palette {
		for _, c := range paletteStr {
			if c > unicode.MaxASCII {
				panic("Palette entry was not ascii")
			}
		}
	}
	for y := 0; y < r.xdim; y++ {
		for z := 0; z < r.ydim; z++ {
			for x := 0; x < r.zdim; x++ {
				if r.Get(x, y, z) < 0 {
					errorMsg := fmt.Sprintf("Block id (%d) was less than zero", r.Get(x, y, z))
					panic(errorMsg)
				}
				if r.Get(x, y, z) >= len(r.palette) {
					errorMsg := fmt.Sprintf("Block id (%d) was greater than or equal to palette length (%d)", r.Get(x, y, z), len(r.palette))
					panic(errorMsg)
				}
			}
		}
	}
}

func (r *Region) IsInRange(x, y, z int) bool {
	return x >= 0 && x < r.xdim && y >= 0 && y < r.ydim && z >= 0 && z < r.zdim
}

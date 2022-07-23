package main

import (
	"fmt"
	"math"

	"github.com/hhhzzzsss/procedura-generator/region"
	"github.com/hhhzzzsss/procedura-generator/util"
)

type blockCacheEntry struct {
	id        int
	pointTrap float64
}

func main() {
	dim := 256
	r := region.MakeRegion(dim, dim, dim)

	r.AddPaletteBlock("air")
	r.AddPaletteBlock("cyan_concrete")
	r.AddPaletteBlock("light_blue_concrete")
	r.AddPaletteBlock("blue_concrete")
	r.AddPaletteBlock("sea_lantern")

	fmt.Println("Calculating mandelbulb...")
	scale := 1.1
	resolution := 0.5 * scale / float64(dim)
	cache := region.MakeRegionCache[blockCacheEntry](&r)
	cache.ForEachNormalizedParallel(8, func(x, y, z float64) blockCacheEntry {
		Z := util.MakeTriplex(0, 0, 0)
		C := util.MakeTriplex(x, y, z).Multiply(1.2)
		// P := util.MakeTriplex(0, 0, 0)
		dZLen := 0.
		ZLen := Z.Length()
		minXDist := 10.0
		minYDist := 10.0
		minZDist := 10.0
		minPDist := 10.0
		for i := 0; i < 12; i++ {
			dZLen = 8*ZLen*ZLen*ZLen*ZLen*ZLen*ZLen*ZLen*dZLen + 1
			Z = Z.Pow(8).Add(C)
			ZLen = Z.Length()
			if ZLen > 256. || dZLen > 1e50 {
				break
			}
			if i < 4 {
				if math.Abs(Z.X) < minXDist {
					minXDist = math.Abs(Z.X)
				}
				if math.Abs(Z.Y) < minYDist {
					minYDist = math.Abs(Z.Y)
				}
				if math.Abs(Z.Z) < minZDist {
					minZDist = math.Abs(Z.Z)
				}
			}
			if i > 0 && i < 5 {
				pDist := (math.Sqrt(math.Abs(Z.X)*math.Abs(Z.X)+math.Abs(Z.Y)*math.Abs(Z.Y)) - 0.05) / dZLen
				if pDist < minPDist {
					minPDist = pDist
				}
			}
		}
		dist := math.Log(ZLen) * ZLen / dZLen
		if dist > resolution {
			return blockCacheEntry{0, 0}
		}
		color := 1
		minDist := minXDist
		if minYDist < minDist {
			minDist = minYDist
			color = 2
		}
		if minZDist < minDist {
			color = 3
		}
		return blockCacheEntry{color, minPDist}
	})

	fmt.Println("Transferring block ids to region")
	r.ForEach(func(x, y, z int) int {
		return cache.Get(x, y, z).id
	})

	fmt.Println("Populating interior of mandelbulb")
	r.ForEachInInterior(func(x, y, z int) {
		if cache.Get(x, y, z).pointTrap < resolution {
			r.Set(x, y, z, 4)
		} else {
			r.Set(x, y, z, 0)
		}
	})

	r.CreateDump()
}

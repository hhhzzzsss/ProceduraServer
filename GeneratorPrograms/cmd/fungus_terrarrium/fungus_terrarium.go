package main

import (
	"flag"
	"math"

	"github.com/hhhzzzsss/procedura-generator/region"
	"github.com/hhhzzzsss/procedura-generator/util"
)

const scale = -1.5
const secondaryScale = 1.0
const outerRadius = 1.0
const innerRadius = 0.5

const outerRadius2 = outerRadius * outerRadius
const innerRadius2 = innerRadius * innerRadius

func main() {
	dim := 256
	r := region.MakeRegion(dim, dim, dim)

	r.AddPaletteBlock("air")
	r.AddPaletteBlock("green_concrete")
	r.AddPaletteBlock("glowstone")
	r.AddPaletteBlock("coarse_dirt")
	r.AddPaletteBlock("warped_hyphae")
	r.AddPaletteBlock("warped_wart_block")
	r.AddPaletteBlock("lime_stained_glass")

	xPtr := flag.Float64("x", -0.236352149685061202, "x position")
	yPtr := flag.Float64("y", 1.60884003225537981, "y position")
	zPtr := flag.Float64("z", -1.5072373200834573, "z position")
	zoomPtr := flag.Float64("zoom", 0.030000000000000006, "zoom factor")
	flag.Parse()

	r.ForEachNormalized(func(x, y, z float64) int {
		v := util.MakeVec3d(x, y, z)
		v = v.Scale(*zoomPtr).Add(util.MakeVec3d(*xPtr, *yPtr, *zPtr))
		c := v
		dr := 1.0
		boxFoldCount := 0
		sphereInvCount := 0
		sphereScaleCount := 0
		for i := 0; i < 256; i++ {
			boxFold(&v, &dr, &boxFoldCount)
			v = v.Scale(secondaryScale)
			dr *= math.Abs(secondaryScale)
			sphereFold(&v, &dr, &sphereInvCount, &sphereScaleCount)
			v = v.Scale(scale).Add(c)
			dr = dr*math.Abs(scale) + 1

			if v.LengthSquared() > 64*64 || dr > 1e12 {
				break
			}
		}
		dist := v.Length() / dr
		if dist < *zoomPtr/float64(dim) {
			boxFoldCount *= 1
			sphereInvCount *= 9
			sphereScaleCount *= 50
			max := boxFoldCount
			if sphereInvCount > max {
				max = sphereInvCount
			}
			if sphereScaleCount > max {
				max = sphereScaleCount
			}
			if max == boxFoldCount {
				return 1
			} else if max == sphereInvCount {
				if sphereInvCount > boxFoldCount*2 {
					return 2
				} else {
					return 3
				}
			} else {
				if sphereScaleCount > boxFoldCount*4 {
					return 4
				} else if sphereScaleCount > boxFoldCount*2 {
					return 5
				} else {
					return 6
				}
			}
		} else {
			return 0
		}
	})

	r.ForEachOnBorder(func(x, y, z int) {
		if r.Get(x, y, z) == 2 {
			r.Set(x, y, z, 3)
		}
	})

	r.CreateDump()
}

func boxFold(v *util.Vec3d, dr *float64, foldCount *int) {
	if v.X > 1 {
		v.X = 2 - v.X
		*foldCount++
	} else if v.X < -1 {
		v.X = -2 - v.X
		*foldCount++
	}
	if v.Y > 1 {
		v.Y = 2 - v.Y
		*foldCount++
	} else if v.Y < -1 {
		v.Y = -2 - v.Y
		*foldCount++
	}
	if v.Z > 1 {
		v.Z = 2 - v.Z
		*foldCount++
	} else if v.Z < -1 {
		v.Z = -2 - v.Z
		*foldCount++
	}
}

func sphereFold(v *util.Vec3d, dr *float64, inversionCount, scaleCount *int) {
	r2 := v.LengthSquared()
	if r2 < innerRadius2 { // Scales inner sphere to outer sphere
		t := outerRadius2 / innerRadius2
		*v = v.Scale(t)
		(*dr) *= t
		*scaleCount++
	} else if r2 < outerRadius { // sphere inversion
		t := outerRadius2 / r2
		*v = v.Scale(t)
		(*dr) *= t
		*inversionCount++
	}
}

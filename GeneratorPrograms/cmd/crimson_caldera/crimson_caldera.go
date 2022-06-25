package main

import (
	"flag"
	"math"

	"github.com/hhhzzzsss/procedura-generator/region"
	"github.com/hhhzzzsss/procedura-generator/util"
)

const scale = 2.25
const secondaryScale = 1.0
const outerRadius = 1.5
const innerRadius = 0.5

const outerRadius2 = outerRadius * outerRadius
const innerRadius2 = innerRadius * innerRadius

func main() {
	dim := 256
	r := region.MakeRegion(dim, dim, dim)

	r.AddPaletteBlock("air")
	r.AddPaletteBlock("smooth_sandstone")
	r.AddPaletteBlock("red_concrete")
	r.AddPaletteBlock("sea_lantern")
	r.AddPaletteBlock("magma_block")
	// r.AddPaletteBlock("red_stained_glass")

	xPtr := flag.Float64("x", 0, "x position")
	yPtr := flag.Float64("y", 0, "y position")
	zPtr := flag.Float64("z", 0, "z position")
	zoomPtr := flag.Float64("zoom", 8, "zoom factor")
	flag.Parse()

	r.ForEachNormalizedParallel(8, func(x, y, z float64) int {
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
			boxFoldCountF := float64(boxFoldCount)
			sphereInvCountF := float64(sphereInvCount)
			sphereScaleCountF := float64(sphereScaleCount)
			boxFoldCountF *= 1
			sphereInvCountF *= 18
			sphereScaleCountF *= 64
			max := boxFoldCountF
			if sphereInvCountF > max {
				max = sphereInvCountF
			}
			if sphereScaleCountF > max {
				max = sphereScaleCountF
			}
			if max == boxFoldCountF {
				return 1
			} else if max == sphereInvCountF {
				if sphereInvCountF > boxFoldCountF*1.1 {
					return 2
				} else {
					return 3
				}
			} else {
				// if sphereScaleCountF > boxFoldCountF*2 {
				// 	return 4
				// } else {
				// 	return 5
				// }
				return 4
			}
		} else {
			return 0
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

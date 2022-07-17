package main

import (
	"flag"
	"math"

	"github.com/hhhzzzsss/procedura-generator/region"
	"github.com/hhhzzzsss/procedura-generator/util"
)

// Pos = 0.49054080746182827 1.813494639229442 0.5298026418567763, Zoom = 0.30000000000000004
// Pos = 0.31320769406691127 1.8134939951797576034435 0.37655575765272236, Zoom = 0.30000000000000004

const scale = -2.5
const secondaryScale = 1.0
const outerRadius = 1.0
const innerRadius = 0.5

const outerRadius2 = outerRadius * outerRadius
const innerRadius2 = innerRadius * innerRadius

func main() {
	minDim := 256
	r := region.MakeRegion(512, 256, 512)

	r.AddPaletteBlock("air")                    // 0
	r.AddPaletteBlock("smooth_stone")           // 1
	r.AddPaletteBlock("glowstone")              // 2
	r.AddPaletteBlock("dirt")                   // 3
	r.AddPaletteBlock("coarse_dirt")            // 4
	r.AddPaletteBlock("cobblestone")            // 5
	r.AddPaletteBlock("stone")                  // 6
	r.AddPaletteBlock("waxed_copper_block")     // 7
	r.AddPaletteBlock("waxed_exposed_copper")   // 8
	r.AddPaletteBlock("waxed_weathered_copper") // 9
	r.AddPaletteBlock("waxed_oxidized_copper")  // 10

	xPtr := flag.Float64("x", 0, "x position")
	yPtr := flag.Float64("y", 0, "y position")
	zPtr := flag.Float64("z", 0, "z position")
	zoomPtr := flag.Float64("zoom", 3, "zoom factor")
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
		if dist < *zoomPtr/float64(minDim) {
			boxFoldCountF := float64(boxFoldCount)
			sphereInvCountF := float64(sphereInvCount)
			sphereScaleCountF := float64(sphereScaleCount)
			boxFoldCountF *= 1
			sphereInvCountF *= 10
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
				if sphereInvCountF > 2.5*boxFoldCountF {
					return 2
				} else if sphereInvCountF > 2.0*boxFoldCountF {
					return 3
				} else if sphereInvCountF > 1.5*boxFoldCountF {
					return 4
				} else if sphereInvCountF > 1.4*boxFoldCountF {
					return 5
				} else {
					return 6
				}
			} else {
				// if sphereInvCountF > 1.5*boxFoldCountF {
				// 	return 2
				// } else if sphereInvCountF > 1.0*boxFoldCountF {
				// 	return 7
				// } else if sphereInvCountF > 0.8*boxFoldCountF {
				// 	return 8
				// } else if sphereInvCountF > 0.4*boxFoldCountF {
				// 	return 9
				// } else {
				// 	return 10
				// }
				if sphereScaleCountF > 5*boxFoldCountF {
					return 7
				} else if sphereScaleCountF > 4*boxFoldCountF {
					return 8
				} else if sphereScaleCountF > 3*boxFoldCountF {
					return 9
				} else {
					return 10
				}
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

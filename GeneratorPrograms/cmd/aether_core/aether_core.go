package main

import (
	"math"

	"github.com/hhhzzzsss/procedura-generator/region"
	"github.com/hhhzzzsss/procedura-generator/util"
)

const zoom = 4.2
const scale = 3.0
const secondaryScale = 1.0
const outerRadius = 1.0
const innerRadius = 0.5

const outerRadius2 = outerRadius * outerRadius
const innerRadius2 = innerRadius * innerRadius

const weathered = false

func main() {
	if weathered {
		weatheredVersion()
	} else {
		cleanVersion()
	}
}

func cleanVersion() {
	dim := 256
	r := region.MakeRegion(dim, dim, dim)

	r.AddPaletteBlock("air")
	r.AddPaletteBlock("smooth_quartz")
	r.AddPaletteBlock("sea_lantern")
	r.AddPaletteBlock("dark_prismarine")
	r.AddPaletteBlock("blue_stained_glass")

	r.ForEachNormalized(func(x, y, z float64) int {
		v := util.MakeVec3d(x, y, z)
		v = v.Scale(zoom)
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
		if dist < zoom/float64(dim) {
			boxFoldCount *= 1
			sphereInvCount *= 10
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
				return 2
			} else {
				return 4
			}
		} else {
			return 0
		}
	})

	r.CreateDump()
}

func weatheredVersion() {
	dim := 256
	r := region.MakeRegion(dim, dim, dim)

	r.AddPaletteBlock("air")
	r.AddPaletteBlock("smooth_quartz")
	r.AddPaletteBlock("sea_lantern")
	r.AddPaletteBlock("dark_prismarine")
	r.AddPaletteBlock("blue_stained_glass")

	r.ForEachNormalized(func(x, y, z float64) int {
		v := util.MakeVec3d(x, y, z)
		v = v.Scale(zoom)
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
		if dist < zoom/float64(dim) {
			boxFoldCount *= 1
			sphereInvCount *= 20
			sphereScaleCount *= 100
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
				if sphereInvCount > 2*boxFoldCount {
					return 2
				} else {
					return 3
				}
			} else {
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

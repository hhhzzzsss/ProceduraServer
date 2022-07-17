package main

import (
	"math"

	"github.com/hhhzzzsss/procedura-generator/region"
	"github.com/hhhzzzsss/procedura-generator/util"
)

func main() {
	dim := 256
	r := region.MakeRegion(dim, dim, dim)

	transparent := false
	if transparent {
		r.AddPaletteBlock("air")
		r.AddPaletteBlock("black_stained_glass")
		r.AddPaletteBlock("white_stained_glass")
		r.AddPaletteBlock("purple_stained_glass")
	} else {
		r.AddPaletteBlock("air")
		r.AddPaletteBlock("smooth_basalt")
		r.AddPaletteBlock("calcite")
		r.AddPaletteBlock("amethyst_block")
	}

	r.ForEachNormalized(func(x, y, z float64) int {
		if y > 0 {
			return 0
		}
		scale := 1.1
		resolution := 2. * scale / float64(dim)
		x = x * scale
		y = y * scale
		z = z*scale + 0.2
		Z := util.MakeQuaternion(x, y, z, 0)
		C := util.MakeQuaternion(-2, 6, 15, -6).Scale(1. / 22.) // C parameter taken from https://www.shadertoy.com/view/3tsyzl
		dZLen := 1.
		ZLen2 := Z.LengthSquared()
		escapeTime := 256
		for i := 0; i < 256; i++ {
			dZLen *= 3 * Z.LengthSquared()
			Z = Z.Cube().Add(C)
			ZLen2 = Z.LengthSquared()
			if ZLen2 > 256. || dZLen > 1e50 {
				escapeTime = i
				break
			}
		}
		dist := 0.5 * math.Log(ZLen2) * math.Sqrt(ZLen2) / dZLen
		if dist < resolution {
			if escapeTime < 30 {
				return 1
			} else if escapeTime < 50 {
				return 2
			} else if escapeTime < 70 {
				return 3
			} else {
				return 0
			}
		} else {
			return 0
		}
	})

	r.CreateDump()
}

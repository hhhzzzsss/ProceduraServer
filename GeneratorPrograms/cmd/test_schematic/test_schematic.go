package main

import (
	"math"

	"github.com/hhhzzzsss/procedura-generator/region"
)

func main() {
	r := region.MakeRegion(256, 256, 256)

	r.AddPaletteBlock("air")
	r.AddPaletteBlock("white_concrete")
	r.AddPaletteBlock("light_blue_stained_glass")
	r.AddPaletteBlock("blue_stained_glass")

	r.ForEachNormalized(func(x, y, z float64) int {
		if x*x+y*y+z*z > 1 {
			return 0
		} else if y < 0.2*math.Sin(x*16)+0.2*math.Sin(z*16) {
			return 1
		} else if y < 0.4*math.Sin(x*16)+0.4*math.Sin(z*16)+0.45 {
			return 2
		} else {
			return 3
		}
	})

	r.CreateDump()
}

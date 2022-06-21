package main

import (
	"flag"
	"math"

	"github.com/hhhzzzsss/procedura-generator/region"
	"github.com/hhhzzzsss/procedura-generator/util"
)

func main() {
	r := region.MakeRegion(256, 256, 256)

	r.AddPaletteBlock("air")
	r.AddPaletteBlock("cyan_concrete")
	r.AddPaletteBlock("light_blue_concrete")
	r.AddPaletteBlock("blue_concrete")

	xPtr := flag.Float64("x", 0.0, "x position")
	yPtr := flag.Float64("y", 0.0, "y position")
	zPtr := flag.Float64("z", 0.0, "z position")
	zoomPtr := flag.Float64("zoom", 1.0, "zoom factor")
	flag.Parse()

	r.ForEachNormalized(func(x, y, z float64) int {
		Z := util.Triplex{}
		C := util.MakeTriplex(x, y, z).Multiply(*zoomPtr).Add(util.MakeTriplex(*xPtr, *yPtr, *zPtr))
		minXDist := 10.0
		minYDist := 10.0
		minZDist := 10.0
		for i := 0; i < 12; i++ {
			Z = Z.Pow(8).Add(C)
			if math.Abs(Z.X) < minXDist {
				minXDist = math.Abs(Z.X)
			}
			if math.Abs(Z.Y) < minYDist {
				minYDist = math.Abs(Z.Y)
			}
			if math.Abs(Z.Z) < minZDist {
				minZDist = math.Abs(Z.Z)
			}
			if Z.LengthSquared() >= 4.0 {
				return 0
			}
		}
		color := 1
		minDist := minXDist
		if minYDist < minDist {
			minDist = minYDist
			color = 2
		}
		if minZDist < minDist {
			// minDist = minZDist
			color = 3
		}
		return color
	})

	r.Hollow()

	r.CreateDump()
}

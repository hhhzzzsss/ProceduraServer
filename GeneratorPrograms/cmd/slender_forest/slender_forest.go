package main

import (
	"fmt"
	"math/rand"

	"github.com/fogleman/poissondisc"
	"github.com/hhhzzzsss/procedura-generator/region"
	"github.com/hhhzzzsss/procedura-generator/treegen"
)

func main() {
	xdim := 256
	ydim := 64
	zdim := 256
	r := region.MakeRegion(xdim, ydim, zdim)

	r.AddPaletteBlock("air")
	r.AddPaletteBlock("grass_block[snowy=true]")
	r.AddPaletteBlock("calcite")
	r.AddPaletteBlock("acacia_wood")

	for x := 0; x < 256; x++ {
		for z := 0; z < 256; z++ {
			r.Set(x, 0, z, 1)
		}
	}

	fmt.Println("Loading roots...")
	roots := make([]*treegen.SkeletonNode, 0)
	poissonPoints := poissondisc.Sample(3, 3, float64(xdim)-3, float64(zdim)-3, 5, 30, nil)
	for _, point := range poissonPoints {
		roots = append(roots, treegen.NewSkeletonNode(point.X, 0.5, point.Y))
	}

	fmt.Println("Loading attractor set...")
	numAttractors := 10000
	attractors := make([]*treegen.Attractor, 0, numAttractors)
	for len(attractors) < numAttractors {
		x := rand.Float64() * float64(xdim)
		y := rand.Float64()*(float64(ydim)-16) + 16
		z := rand.Float64() * float64(zdim)
		attractors = append(attractors, treegen.NewAttractor(x, y, z))
	}

	fmt.Println("Generating skeleton")
	settings := treegen.GetDefaultSettings()
	settings.StepSize = 0.5
	settings.KillDistance = 1
	settings.AttractionRadius = 20
	settings.BaseThickness = 1
	skeleton := treegen.GenerateSkeleton(roots, attractors, settings)

	fmt.Println("Fleshing out branches from skeleton...")
	skeleton.ForEachNode(func(node *treegen.SkeletonNode) {
		if node.IsRoot() {
			return
		}
		x := node.GetDim(0)
		y := node.GetDim(1)
		z := node.GetDim(2)
		bx := int(x)
		by := int(y)
		bz := int(z)
		r.Set(bx, by, bz, 3)
		r.ForEachInSphere(x, y, z, node.GetThickness(), func(sx, sy, sz int, rad2 float64) {
			r.Set(sx, sy, sz, 3)
		})
	})

	r.CreateDump()
}

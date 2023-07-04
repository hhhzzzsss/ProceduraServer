package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/hhhzzzsss/procedura-generator/region"
	"github.com/hhhzzzsss/procedura-generator/treegen"
	"github.com/hhhzzzsss/procedura-generator/util"
	"github.com/ojrac/opensimplex-go"
)

const dim = 256
const halfDim = float64(dim) / 2.

func volumeNoise(noise opensimplex.Noise, p util.Vec3d) float64 {
	return noise.Eval3(p.X, p.Y, p.Z)
}

func fractalNoise(noise opensimplex.Noise, p util.Vec3d, iter int) float64 {
	val := 0.0
	scale := 1.0
	for i := 0; i < iter; i++ {
		val += volumeNoise(noise, p.Scale(scale)) / scale
		scale *= 2.0
	}
	return val
}

func terrainNoise(noise opensimplex.Noise, pNorm util.Vec3d) float64 {
	val := fractalNoise(noise, pNorm, 5)
	val = val*0.1 + 0.01
	return val
}

func treeNoise(noise opensimplex.Noise, pNorm util.Vec3d) float64 {
	val := fractalNoise(noise, pNorm.Add(util.MakeVec3d(492.32, 109.38, -193.04)), 3)
	return val*0.03 + 0.01
}

func broadCaveNoise(noise opensimplex.Noise, p util.Vec3d) float64 {
	p = p.Scale(1).Add(util.MakeVec3d(203.46, 38.19, -592.34))
	val := fractalNoise(noise, p, 4)
	return 0.4 - math.Abs(val)
}

func narrowCaveNoise(noise opensimplex.Noise, p util.Vec3d) util.Vec3d {
	pX := p.InvScale(20).Add(util.MakeVec3d(-864.90, -236.04, -933.49))
	pY := p.InvScale(20).Add(util.MakeVec3d(-550.39, 35.08, 652.52))
	pZ := p.InvScale(20).Add(util.MakeVec3d(-520.13, 316.94, -144.99))
	return util.MakeVec3d(fractalNoise(noise, pX, 3), fractalNoise(noise, pY, 3), fractalNoise(noise, pZ, 3))
}

func perlinWorms(noise opensimplex.Noise, r *region.Region) region.RegionCache[float64] {
	rc := region.MakeRegionCache[float64](r)
	for i := 0; i < 200; i++ {
		p := util.MakeVec3d(rand.Float64()*256, rand.Float64()*256, rand.Float64()*256)
		for i := 0; i < 200; i++ {
			dir := narrowCaveNoise(noise, p)
			dirLen := dir.Length()
			if dirLen < 0.001 {
				break
			}
			p = p.Add(dir.InvScale(dirLen))
			if p.X < 0 || p.X >= 256 || p.Y < 0 || p.Y >= 256 || p.Z < 0 || p.Z >= 256 {
				break
			}
			width := fractalNoise(noise, p.InvScale(16).Add(util.MakeVec3d(-30.98, 440.80, -418.66)), 3)/2.0 + 3
			r.ForEachInSphere(p.X, p.Y, p.Z, width, func(x, y, z int, rad2 float64) {
				rc.Set(x, y, z, math.Max(rc.Get(x, y, z), (width-math.Sqrt(rad2))/3))
			})
		}
	}
	return rc
}

func doDLA(particles []util.Vec3i, r *region.Region) {
	maxIter := 10000
	maxParticles := len(particles)
	rc := region.MakeRegionCache[int](r)
	rc.ForEach(func(x, y, z int) int {
		blockId := r.Get(x, y, z)
		if blockId == 0 {
			return 0 // Air
		} else if blockId == 9 {
			return 2 // Slightly sticky
		} else {
			return 1 // Not sticky
		}
	})
	for i := 0; i < maxIter; i++ {
		fmt.Print("\rParticles left: ", len(particles), "/", maxParticles, "        ")
		for i := 0; i < len(particles); i++ {
			p := particles[i]
			validNeighbors := make([]util.Vec3i, 0)
			maxSticky := 0
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					for dz := -1; dz <= 1; dz++ {
						if dx == 0 && dy == 0 && dz == 0 {
							continue
						}
						neighborCandidate := util.MakeVec3i(p.X+dx, p.Y+dy, p.Z+dz)
						if neighborCandidate.Sub(util.MakeVec3i(int(halfDim), int(halfDim), int(halfDim))).Length() > halfDim*0.7 {
							continue
						}
						neighborType := rc.Get(neighborCandidate.X, neighborCandidate.Y, neighborCandidate.Z)
						if neighborType == 0 {
							validNeighbors = append(validNeighbors, neighborCandidate)
						} else if neighborType > maxSticky {
							maxSticky = neighborType
						}
					}
				}
			}

			stick := false
			if maxSticky == 2 {
				if rand.Float64() < 0.1 {
					stick = true
				}
			} else if maxSticky == 3 {
				stick = true
			}
			if stick {
				rc.Set(p.X, p.Y, p.Z, 3)
				if len(particles) > 9*maxParticles/10 {
					r.Set(p.X, p.Y, p.Z, 10)
				} else if len(particles) > 8*maxParticles/10 {
					r.Set(p.X, p.Y, p.Z, 11)
				} else if len(particles) > 7*maxParticles/10 {
					r.Set(p.X, p.Y, p.Z, 12)
				} else if len(particles) > 5*maxParticles/10 {
					r.Set(p.X, p.Y, p.Z, 6)
				} else if len(particles) > 3*maxParticles/10 {
					r.Set(p.X, p.Y, p.Z, 5)
				} else if len(particles) > 1*maxParticles/10 {
					r.Set(p.X, p.Y, p.Z, 13)
				} else {
					r.Set(p.X, p.Y, p.Z, 14)
				}
				for _, neighborPos := range validNeighbors {
					r.Set(neighborPos.X, neighborPos.Y, neighborPos.Z, 15)
				}
				util.RemoveFromUnorderedSlice[util.Vec3i](&particles, i)
				i--
			} else if len(validNeighbors) > 0 {
				particles[i] = validNeighbors[rand.Int()%len(validNeighbors)]
			} else {
				util.RemoveFromUnorderedSlice[util.Vec3i](&particles, i)
				i--
			}
		}
		if len(particles) == 0 {
			break
		}
	}
	fmt.Println("\rParticles left: ", len(particles), "/", maxParticles)
}

func main() {
	r := region.MakeRegion(dim, dim, dim)

	r.AddPaletteBlock("air")                      // 0
	r.AddPaletteBlock("stone")                    // 1
	r.AddPaletteBlock("coarse_dirt")              // 2
	r.AddPaletteBlock("moss_block")               // 3
	r.AddPaletteBlock("sand")                     // 4
	r.AddPaletteBlock("light_blue_stained_glass") // 5
	r.AddPaletteBlock("blue_stained_glass")       // 6
	r.AddPaletteBlock("oak_wood")                 // 7
	r.AddPaletteBlock("oak_leaves")               // 8
	r.AddPaletteBlock("magma_block")              // 9
	r.AddPaletteBlock("crying_obsidian")          // 10
	r.AddPaletteBlock("purple_concrete")          // 11
	r.AddPaletteBlock("blue_concrete")            // 12
	r.AddPaletteBlock("white_stained_glass")      // 13
	r.AddPaletteBlock("glowstone")                // 14
	r.AddPaletteBlock("light[level=8]")           // 15
	r.AddPaletteBlock("grass_block")              // 16
	r.AddPaletteBlock("dirt")                     // 17
	r.AddPaletteBlock("bedrock")                  // 18

	noise1 := opensimplex.New(1234) // terrain
	noise2 := opensimplex.New(4321) // caves
	noise3 := opensimplex.New(4929) // trees

	seaLevel := 0.7

	rand.Seed(222) // perlin worms seed
	fmt.Println("Generating perlin worms...")
	wormCache := perlinWorms(noise2, &r)

	rand.Seed(1337) // DLA particles seed
	particles := make([]util.Vec3i, 0)
	fmt.Println("Constructing planet...")
	r.ForEachParallel(8, func(bx, by, bz int) int {
		x := 2.0*float64(bx)/float64(dim) - 1.0
		y := 2.0*float64(by)/float64(dim) - 1.0
		z := 2.0*float64(bz)/float64(dim) - 1.0

		p := util.MakeVec3d(x, y, z)
		pLen := p.Length()
		pNorm := p.InvScale(pLen)
		terrainHeight := seaLevel + terrainNoise(noise1, pNorm)

		if pLen > 0.8 && p.Add(util.MakeVec3d(0.0, 0.3, 0.0)).Length() > 1.0 {
			if by >= 50 {
				return 0 // air
			} else if by == 49 {
				return 16 // grass_block
			} else if by >= 17 && by <= 48 {
				return 17 // dirt
			} else if by >= 1 && by <= 16 {
				return 1 // stone
			} else {
				return 18 // bedrock
			}
		}

		// Magma core
		if pLen < 0.15 {
			return 9
		}

		surfaceCloseness := 5.0 / math.Max((terrainHeight-pLen)*128, 1.0)
		broadCaveFactor := broadCaveNoise(noise2, p) - surfaceCloseness
		surfaceCloseness = math.Max(surfaceCloseness-0.4, 0.0)
		waterCloseness := surfaceCloseness * util.Clamp((seaLevel-terrainHeight)*128+1, 0.0, 1.0)
		narrowCaveFactor := wormCache.Get(bx, by, bz) - waterCloseness

		if broadCaveFactor > 0.0 || narrowCaveFactor > 0.0 {
			if pLen < terrainHeight {
				if rand.Float64() < 0.02 {
					particles = append(particles, util.MakeVec3i(bx, by, bz))
				}
			}
			return 0
		}

		// Stone interior
		if pLen < terrainHeight-(5./halfDim) {
			return 1
		}

		if terrainHeight > seaLevel { // Above sea level
			if pLen < terrainHeight-(1./halfDim) {
				return 2
			} else if pLen < terrainHeight {
				return 3
			}
		} else { // Below sea level
			if pLen < terrainHeight {
				return 4
			} else if pLen < seaLevel-(1./halfDim) {
				return 6
			} else if pLen < seaLevel {
				return 5
			}
		}

		return 0
	})

	fmt.Println("Loading trees...")

	rand.Seed(131313)
	rootPlacementAttempts := 1000
	roots := make([]*treegen.SkeletonNode, 0)
outer:
	for i := 0; i < rootPlacementAttempts; i++ {
		x := 2.0*rand.Float64() - 1.0
		y := 2.0*rand.Float64() - 1.0
		z := 2.0*rand.Float64() - 1.0
		p := util.MakeVec3d(x, y, z)
		pLen := p.Length()
		if pLen > 1.0 || pLen < 0.0001 {
			continue
		}
		pNorm := p.InvScale(pLen)
		terrainHeight := seaLevel + terrainNoise(noise1, pNorm)
		if terrainHeight < seaLevel {
			continue
		}
		// waterFallof := (math.Min((terrainHeight-seaLevel)*halfDim, 1.5) - 1.5) / 2.
		treeFactor := treeNoise(noise3, pNorm)
		if treeFactor <= 0.0 {
			continue
		}
		blockPos := pNorm.Scale(terrainHeight * halfDim).Add(util.MakeVec3d(halfDim, halfDim, halfDim))
		if wormCache.Get(int(blockPos.X), int(blockPos.Y), int(blockPos.Z)) > 0.0 {
			continue
		}
		// for r.Get(int(blockPos.X), int(blockPos.Y), int(blockPos.Z)) != 0 {
		// 	blockPos = blockPos.Add(pNorm.Scale(0.2))
		// }
		newRoot := treegen.NewSkeletonNode(blockPos.X, blockPos.Y, blockPos.Z)
		for _, root := range roots {
			if util.PointDistSq(newRoot, root) < 12*12 {
				continue outer
			}
		}

		roots = append(roots, newRoot)
	}

	numAttractors := 20000
	attractors := make([]*treegen.Attractor, 0, numAttractors)
	for len(attractors) < numAttractors {
		x := 2.0*rand.Float64() - 1.0
		y := 2.0*rand.Float64() - 1.0
		z := 2.0*rand.Float64() - 1.0
		p := util.MakeVec3d(x, y, z)
		pLen := p.Length()
		if pLen > 1.0 || pLen < 0.0001 {
			continue
		}
		pNorm := p.InvScale(pLen)
		terrainHeight := seaLevel + terrainNoise(noise1, pNorm)
		if terrainHeight < seaLevel {
			continue
		}
		// waterFallof := (math.Min((terrainHeight-seaLevel)*halfDim, 1.5) - 1.5) / 2.
		treeFactor := treeNoise(noise3, pNorm)
		if rand.Float64() > treeFactor {
			continue
		}
		blockPos := pNorm.Scale(terrainHeight * halfDim).Add(util.MakeVec3d(halfDim, halfDim, halfDim))
		attractors = append(attractors, treegen.NewAttractor(blockPos.X, blockPos.Y, blockPos.Z))
	}

	settings := treegen.GetDefaultSettings()
	settings.StepSize = 0.5
	settings.KillDistance = 1
	settings.AttractionRadius = 10
	settings.DoThicknessPostprocess = false
	skeleton := treegen.GenerateSkeleton(roots, attractors, settings)

	maxThickness := 0.0
	skeleton.ForEachNode(func(node *treegen.SkeletonNode) {
		if node.GetThickness() > maxThickness {
			maxThickness = node.GetThickness()
		}
	})

	skeleton.ForEachNode(func(node *treegen.SkeletonNode) {
		p := util.MakeVec3d(node.GetDim(0), node.GetDim(1), node.GetDim(2))
		center := util.MakeVec3d(halfDim, halfDim, halfDim)
		pNorm := p.Sub(center).InvScale(halfDim)
		for r.Get(int(p.X), int(p.Y), int(p.Z)) != 0 && r.Get(int(p.X), int(p.Y), int(p.Z)) != 7 && r.Get(int(p.X), int(p.Y), int(p.Z)) != 8 {
			p = p.Add(pNorm.Scale(0.1))
		}
		bx := int(p.X)
		by := int(p.Y)
		bz := int(p.Z)
		if node.GetThickness() <= 2.0 {
			r.Set(bx, by, bz, 8)
		} else {
			r.Set(bx, by, bz, 7)
		}
		r.ForEachInSphere(p.X, p.Y, p.Z, 2.5*node.GetThickness()/maxThickness, func(sx, sy, sz int, rad2 float64) {
			r.Set(sx, sy, sz, 7)
		})
	})

	// skeleton.ForEachNode(func(node *treegen.SkeletonNode) {
	// 	x := node.GetDim(0)
	// 	y := node.GetDim(1)
	// 	z := node.GetDim(2)
	// 	if node.GetThickness() == 1.0 {
	// 		r.ForEachInSphere(x, y, z, 2.5*node.GetRoot().GetThickness()/maxThickness, func(sx, sy, sz int, rad2 float64) {
	// 			if r.Get(sx, sy, sz) == 0 {
	// 				r.Set(sx, sy, sz, 8)
	// 			}
	// 		})
	// 	}
	// })

	// featurePlacementAttempts := 1000
	// for i := 0; i < featurePlacementAttempts; i++ {
	// 	x := 2.0*rand.Float64() - 1.0
	// 	y := 2.0*rand.Float64() - 1.0
	// 	z := 2.0*rand.Float64() - 1.0
	// 	p := util.MakeVec3d(x, y, z)
	// 	pLen := p.Length()
	// 	if pLen > 1.0 || pLen < 0.0001 {
	// 		continue
	// 	}
	// 	pNorm := p.InvScale(pLen)
	// 	terrainHeight := seaLevel + terrainNoise(noise1, pNorm)
	// 	if terrainHeight < seaLevel {
	// 		continue
	// 	}
	// 	blockPos := pNorm.Scale(terrainHeight * halfDim).Add(util.MakeVec3d(halfDim, halfDim, halfDim))
	// 	if wormCache.Get(int(blockPos.X), int(blockPos.Y), int(blockPos.Z)) > 0.0 {
	// 		continue
	// 	}
	// 	for i := 0; i < 5; i++ {
	// 		if r.Get(int(blockPos.X), int(blockPos.Y), int(blockPos.Z)) != 0 {
	// 			blockPos = blockPos.Add(pNorm.Scale(0.2))
	// 		}
	// 	}
	// 	if r.Get(int(blockPos.X), int(blockPos.Y), int(blockPos.Z)) != 0 {
	// 		continue
	// 	}
	// 	r.Set(int(blockPos.X), int(blockPos.Y), int(blockPos.Z), 9)
	// }

	fmt.Println("Doing DLA calculation...")
	rand.Seed(123456)
	doDLA(particles, &r)

	fmt.Println("Hollowing out core...")
	r.SelectiveHollow(9)

	r.CreateDump()
}

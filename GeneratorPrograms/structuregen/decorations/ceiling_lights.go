package decorations

import (
	"math/rand"

	"github.com/hhhzzzsss/procedura-generator/structuregen/block"
	"github.com/hhhzzzsss/procedura-generator/structuregen/direction"
	"github.com/hhhzzzsss/procedura-generator/util"
)

func TileCeiling(tile Decoration, tileX, tileZ int, meta DecorationMeta) Decoration {
	if meta.XDim <= 0 || meta.ZDim <= 0 {
		panic("Invalid xdim or ZDim")
	}

	dec := make(Decoration)

	numTilesX := meta.XDim / tileX
	numTilesZ := meta.ZDim / tileZ
	offsetX := (meta.XDim - tileX*numTilesX) / 2
	offsetZ := (meta.ZDim - tileZ*numTilesZ) / 2
	for xidx := 0; xidx < numTilesX; xidx++ {
		for zidx := 0; zidx < numTilesZ; zidx++ {
			tilepos := util.MakeVec3i(offsetX+xidx*tileX, 0, offsetZ+zidx*tileZ)
			for pos, block := range tile {
				dec[tilepos.Add(pos)] = block
			}
		}
	}

	return dec
}

func CeilingLight0(meta DecorationMeta) Decoration {
	if meta.XDim <= 0 || meta.ZDim <= 0 {
		panic("Invalid xdim or ZDim")
	}

	dec := make(Decoration)

	light_block := block.RandBlock(block.LIGHT_BLOCKS)

	wood_material := block.RandMat(block.WOOD_MATERIALS)
	trapdoor_material := wood_material + "_trapdoor"
	west_trapdoor := block.MakeBlock(trapdoor_material, map[string]string{"facing": "west", "half": "top", "open": "true"})
	north_trapdoor := block.MakeBlock(trapdoor_material, map[string]string{"facing": "north", "half": "top", "open": "true"})
	east_trapdoor := block.MakeBlock(trapdoor_material, map[string]string{"facing": "east", "half": "top", "open": "true"})
	south_trapdoor := block.MakeBlock(trapdoor_material, map[string]string{"facing": "south", "half": "top", "open": "true"})
	north_underside_trapdoor := block.MakeBlock(trapdoor_material, map[string]string{"facing": "north", "half": "top"})
	south_underside_trapdoor := block.MakeBlock(trapdoor_material, map[string]string{"facing": "south", "half": "top"})

	var hang_block block.Block
	hang_rand := rand.Intn(5)
	switch hang_rand {
	case 0:
		hang_block = block.MakeBlock(wood_material+"_fence", nil)
	case 1:
		hang_block = block.MakeBlock("iron_bars", nil)
	case 2:
		hang_block = block.MakeBlock("end_rod", map[string]string{"facing": "up"})
	case 3:
		hang_block = block.MakeBlock("chain", map[string]string{"axis": "y"})
	case 4:
		hang_block = block.MakeBlock("lightning_rod", map[string]string{"facing": "down"})
	}

	bottom_cover := rand.Intn(3) == 0
	hang := rand.Intn(3) == 0

	yoffset := 0
	if hang {
		yoffset = -1
	}

	numLamps := meta.ZDim / 4
	zoffset := (meta.ZDim - numLamps*4) / 2
	for i := 0; i < numLamps; i++ {
		for x := 1; x < meta.XDim-1; x++ {
			if hang {
				dec.SetBlock(x, 0, zoffset+4*i+1, hang_block)
				dec.SetBlock(x, 0, zoffset+4*i+2, hang_block)
			}
			dec.SetBlock(x, yoffset, zoffset+4*i+0, north_trapdoor)
			dec.SetBlock(x, yoffset, zoffset+4*i+1, light_block)
			dec.SetBlock(x, yoffset, zoffset+4*i+2, light_block)
			dec.SetBlock(x, yoffset, zoffset+4*i+3, south_trapdoor)
			if bottom_cover {
				dec.SetBlock(x, yoffset-1, zoffset+4*i+1, south_underside_trapdoor)
				dec.SetBlock(x, yoffset-1, zoffset+4*i+2, north_underside_trapdoor)
			}
		}
		dec.SetBlock(0, yoffset, zoffset+4*i+1, west_trapdoor)
		dec.SetBlock(0, yoffset, zoffset+4*i+2, west_trapdoor)
		dec.SetBlock(meta.XDim-1, yoffset, zoffset+4*i+1, east_trapdoor)
		dec.SetBlock(meta.XDim-1, yoffset, zoffset+4*i+2, east_trapdoor)
	}

	return dec
}

func CeilingLight1(meta DecorationMeta) Decoration {
	tile := make(Decoration)

	surrounding_signs := rand.Intn(3) > 0
	hanging := rand.Intn(3) > 0
	hopper_base := rand.Intn(3) == 0
	trapdoor_fan := surrounding_signs && hanging && rand.Intn(2) == 0 // Requires hanging, and looks kinda empty to have a trapdoor fan without surrounding signs

	lantern := block.MakeBlock("lantern", map[string]string{"hanging": "true"})

	wood_material := block.RandMat(block.WOOD_MATERIALS)
	sign_material := wood_material + "_wall_sign"
	west_sign := block.MakeBlock(sign_material, map[string]string{"facing": "west"})
	north_sign := block.MakeBlock(sign_material, map[string]string{"facing": "north"})
	east_sign := block.MakeBlock(sign_material, map[string]string{"facing": "east"})
	south_sign := block.MakeBlock(sign_material, map[string]string{"facing": "south"})

	hopper := block.MakeBlock("hopper", nil)

	trapdoor_material := wood_material + "_trapdoor"
	west_trapdoor := block.MakeBlock(trapdoor_material, map[string]string{"facing": "west"})
	north_trapdoor := block.MakeBlock(trapdoor_material, map[string]string{"facing": "north"})
	east_trapdoor := block.MakeBlock(trapdoor_material, map[string]string{"facing": "east"})
	south_trapdoor := block.MakeBlock(trapdoor_material, map[string]string{"facing": "south"})

	var chain block.Block
	chain_rand := rand.Intn(6)
	switch chain_rand {
	case 0:
		chain = block.MakeBlock(wood_material+"_fence", nil)
	case 1:
		chain = block.MakeBlock("iron_bars", nil)
	case 2:
		chain = block.MakeBlock("end_rod", map[string]string{"facing": "up"})
	case 3:
		chain = block.MakeBlock("chain", map[string]string{"axis": "y"})
	case 4:
		chain = block.MakeBlock("lightning_rod", map[string]string{"facing": "up"})
	case 5:
		chain = block.MakeBlock(block.RandMat(block.GRAYSCALE_COLORS)+"_stained_glass_pane", nil)
	}

	// black stained glass goes well with hoppers
	if hopper_base && rand.Intn(3) == 0 {
		chain = block.MakeBlock("black_stained_glass_pane", nil)
	}

	xOffset := 1
	zOffset := 1
	if trapdoor_fan {
		xOffset = 2
		zOffset = 2
	}
	yOffset := 0
	if hanging {
		yOffset -= 1
	}
	if hopper_base {
		yOffset -= 1
	}

	tile.SetBlock(xOffset, yOffset, zOffset, lantern)
	if surrounding_signs {
		tile.SetBlock(xOffset-1, yOffset, zOffset+0, west_sign)
		tile.SetBlock(xOffset+0, yOffset, zOffset-1, north_sign)
		tile.SetBlock(xOffset+1, yOffset, zOffset+0, east_sign)
		tile.SetBlock(xOffset+0, yOffset, zOffset+1, south_sign)
	}
	if hanging {
		tile.SetBlock(xOffset, yOffset+1, zOffset, chain)
	}
	if trapdoor_fan {
		tile.SetBlock(xOffset-1, yOffset+1, zOffset+0, west_trapdoor)
		tile.SetBlock(xOffset+0, yOffset+1, zOffset-1, north_trapdoor)
		tile.SetBlock(xOffset+1, yOffset+1, zOffset+0, east_trapdoor)
		tile.SetBlock(xOffset+0, yOffset+1, zOffset+1, south_trapdoor)
	}
	if hopper_base {
		tile.SetBlock(xOffset, 0, zOffset, hopper)
	}

	if trapdoor_fan {
		return TileCeiling(tile, 5, 5, meta)
	} else {
		return TileCeiling(tile, 3, 3, meta)
	}
}

func CeilingLight2(meta DecorationMeta) Decoration {
	tile := make(Decoration)

	var light block.Block
	light_rand := rand.Intn(6)
	switch light_rand {
	case 0:
		light = block.MakeBlock("glowstone", nil)
	case 1:
		light = block.MakeBlock("sea_lantern", nil)
	case 2:
		light = block.MakeBlock("shroomlight", nil)
	case 3:
		light = block.MakeBlock("beacon", nil)
	case 4:
		light = block.MakeBlock("lantern", map[string]string{"hanging": "true"})
	case 5:
		light = block.MakeBlock("end_rod", map[string]string{"facing": "down"})
	}

	var hang_material string
	hang_rand := rand.Intn(3)
	switch hang_rand {
	case 0:
		hang_material = block.RandMat(block.WOOD_MATERIALS) + "_fence"
	case 1:
		hang_material = block.RandMat(block.DYE_COLORS) + "_stained_glass_pane"
	case 2:
		hang_material = "iron_bars"
	}

	top_hang := block.MakeBlock(hang_material, nil)
	center_hang := block.MakeBlock(hang_material, map[string]string{"west": "true", "north": "true", "east": "true", "south": "true"})

	if rand.Intn(3) == 0 {
		top_hang = block.MakeBlock("chain", map[string]string{"axis": "y"})
	}

	tile.SetBlock(2, 0, 2, top_hang)
	tile.SetBlock(2, -1, 2, center_hang)
	for i := 0; i < 4; i++ {
		side_hang := block.MakeBlock(hang_material, map[string]string{direction.DirectionNames[(i+2)%4]: "true"})
		offset := direction.DirectionOffsets[i]
		tile.SetBlock(2+offset.X, -1, 2+offset.Z, side_hang)
		tile.SetBlock(2+offset.X, -2, 2+offset.Z, light)
	}

	return TileCeiling(tile, 5, 5, meta)
}

var ceilingGenerators = []DecorationGenerator{
	CeilingLight0,
	CeilingLight1,
	CeilingLight2,
}

func RandomCeilingLight(meta DecorationMeta) Decoration {
	generator := ceilingGenerators[rand.Intn(len(ceilingGenerators))]
	return generator(meta)
}

package decorations

import (
	"math/rand"

	"github.com/hhhzzzsss/procedura-generator/structuregen/block"
)

func Chair0(meta DecorationMeta) Decoration {
	dec := make(Decoration)

	width := meta.ZDim
	wood_material := block.RandMat(block.WOOD_MATERIALS)
	var stair_material string
	if rand.Intn(2) == 0 {
		stair_material = block.RandMat(block.SLAB_MATERIALS)
	} else {
		stair_material = wood_material
	}

	stair := block.MakeBlock(stair_material+"_stairs", map[string]string{"facing": "west", "half": "bottom"})

	for z := 0; z < width; z++ {
		dec.SetBlock(0, 0, z, stair)
	}

	if meta.IncludeArmrests {
		var north_armrest, south_armrest block.Block
		if rand.Intn(2) == 0 {
			north_armrest = block.MakeBlock(wood_material+"_trapdoor", map[string]string{"facing": "north", "open": "true"})
			south_armrest = block.MakeBlock(wood_material+"_trapdoor", map[string]string{"facing": "south", "open": "true"})
		} else {
			north_armrest = block.MakeBlock(wood_material+"_wall_sign", map[string]string{"facing": "north"})
			south_armrest = block.MakeBlock(wood_material+"_wall_sign", map[string]string{"facing": "south"})
		}
		dec.SetBlock(0, 0, -1, north_armrest)
		dec.SetBlock(0, 0, width, south_armrest)
	}

	return dec
}

func Chair1(meta DecorationMeta) Decoration {
	dec := make(Decoration)

	width := meta.ZDim
	wood_material := block.RandMat(block.WOOD_MATERIALS)
	dye_color := block.RandMat(block.DYE_COLORS)
	var slab_material string
	if rand.Intn(2) == 0 {
		slab_material = block.RandMat(block.SLAB_MATERIALS)
	} else {
		slab_material = wood_material
	}

	var seat block.Block
	if rand.Intn(2) == 0 {
		seat = block.MakeBlock(slab_material+"_slab", map[string]string{"type": "bottom"})
	} else {
		seat = block.MakeBlock(dye_color+"_bed", map[string]string{"facing": "west", "part": "foot"})
	}

	var bottom_back, top_back block.Block
	if rand.Intn(2) == 0 {
		var hinge_side string
		if rand.Intn(2) == 0 {
			hinge_side = "left"
		} else {
			hinge_side = "right"
		}
		bottom_back = block.MakeBlock(wood_material+"_door", map[string]string{"facing": "west", "half": "lower", "hinge": hinge_side})
		top_back = block.MakeBlock(wood_material+"_door", map[string]string{"facing": "west", "half": "upper", "hinge": hinge_side})
	} else {
		bottom_back = block.MakeBlock(wood_material+"_trapdoor", map[string]string{"facing": "west", "half": "bottom", "open": "true"})
		top_back = block.MakeBlock(wood_material+"_trapdoor", map[string]string{"facing": "west", "half": "top", "open": "true"})
	}

	banner := block.MakeBlock(dye_color+"_wall_banner", map[string]string{"facing": "east"})

	for z := 0; z < width; z++ {
		dec.SetBlock(0, 0, z, seat)
		dec.SetBlock(-1, 0, z, bottom_back)
		dec.SetBlock(-1, 1, z, top_back)
		dec.SetBlock(0, 1, z, banner)
	}

	if meta.IncludeArmrests {
		var north_armrest, south_armrest block.Block
		if rand.Intn(2) == 0 {
			north_armrest = block.MakeBlock(wood_material+"_trapdoor", map[string]string{"facing": "north", "open": "true"})
			south_armrest = block.MakeBlock(wood_material+"_trapdoor", map[string]string{"facing": "south", "open": "true"})
		} else {
			north_armrest = block.MakeBlock(wood_material+"_wall_sign", map[string]string{"facing": "north"})
			south_armrest = block.MakeBlock(wood_material+"_wall_sign", map[string]string{"facing": "south"})
		}
		dec.SetBlock(0, 0, -1, north_armrest)
		dec.SetBlock(0, 0, width, south_armrest)
	}

	return dec
}

var chairGenerators = []DecorationGenerator{
	Chair0,
	Chair1,
}

func RandomChair(meta DecorationMeta) Decoration {
	generator := chairGenerators[rand.Intn(len(chairGenerators))]
	return generator(meta)
}

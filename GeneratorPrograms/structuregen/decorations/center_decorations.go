package decorations

import (
	"math/rand"

	"github.com/hhhzzzsss/procedura-generator/structuregen/block"
	"github.com/hhhzzzsss/procedura-generator/structuregen/direction"
	"github.com/hhhzzzsss/procedura-generator/util"
)

func lineChairs(dec, chair *Decoration, x, y, z, rotation, spacing, length int) {
	rotatedChair := chair.Rotate(rotation)

	chairInterval := spacing + 1
	numChairs := (length + spacing) / chairInterval
	startingOffset := ((length + spacing) - numChairs*chairInterval + rand.Intn(2)) / 2

	vdir := direction.DirectionOffsets[(rotation+3)%4]
	startingPos := util.MakeVec3i(x, y, z).Add(vdir.Scale(startingOffset))
	for i := 0; i < numChairs; i++ {
		chairPos := startingPos.Add(vdir.Scale(chairInterval * i))
		for pos, block := range rotatedChair {
			(*dec)[chairPos.Add(pos)] = block
		}
	}
}

func makeBorder(dec *Decoration, x1, z1, x2, z2 int) {
	for x := x1; x <= x2; x++ {
		dec.SetBlock(x, 0, z1, block.AIR)
		dec.SetBlock(x, 0, z2, block.AIR)
	}
	for z := z1 + 1; z <= z2-1; z++ {
		dec.SetBlock(x1, 0, z, block.AIR)
		dec.SetBlock(x2, 0, z, block.AIR)
	}
}

// Table
func CenterDecoration0(meta DecorationMeta) Decoration {
	dec := make(Decoration)

	full_carpet := rand.Intn(5) == 0

	wood_material := block.RandMat(block.WOOD_MATERIALS)
	fence := block.MakeBlock(wood_material+"_fence", nil)
	carpet := block.MakeBlock(block.RandMat(block.DYE_COLORS)+"_carpet", nil)

	makeBorder(&dec, -1, -1, 3, 4)

	for x := 0; x < 3; x++ {
		for z := 0; z < 4; z++ {
			dec.SetBlock(x, 0, z, block.BARRIER)
			dec.SetBlock(x, 1, z, carpet)
		}
	}
	dec.SetBlock(0, 0, 0, fence)
	dec.SetBlock(2, 0, 0, fence)
	dec.SetBlock(0, 0, 3, fence)
	dec.SetBlock(2, 0, 3, fence)

	if !full_carpet {
		north_trapdoor := block.MakeBlock(wood_material+"_trapdoor", map[string]string{"facing": "north", "half": "top"})
		south_trapdoor := block.MakeBlock(wood_material+"_trapdoor", map[string]string{"facing": "south", "half": "top"})
		dec.SetBlock(1, 0, 1, south_trapdoor)
		dec.SetBlock(1, 0, 2, north_trapdoor)
		dec.SetBlock(1, 1, 1, GetOneBlockDecoration(0.15))
		dec.SetBlock(1, 1, 2, GetOneBlockDecoration(0.15))
	}

	switch rand.Intn(2) {
	case 0:
		chair := RandomChair(DecorationMeta{ZDim: 1, IncludeArmrests: true})
		if rand.Intn(2) == 0 {
			lineChairs(&dec, &chair, -1, 0, 1, 0, 2, 2)
			lineChairs(&dec, &chair, 3, 0, 2, 2, 2, 2)
		}
		if rand.Intn(2) == 0 {
			dec.ApplyDecoration(1, 0, -1, chair.Rotate(1))
			dec.ApplyDecoration(1, 0, 4, chair.Rotate(3))
		}
	case 1:
		chair := RandomChair(DecorationMeta{ZDim: 2, IncludeArmrests: true})
		dec.ApplyDecoration(-1, 0, 1, chair)
		dec.ApplyDecoration(3, 0, 2, chair.Rotate(2))
	}

	return dec
}

// Coral Carpet
func CenterDecoration1(meta DecorationMeta) Decoration {
	dec := make(Decoration)

	var coral block.Block
	var floor block.Block
	switch rand.Intn(3) {
	case 0:
		coral = block.MakeBlock("brain_coral_fan", map[string]string{"waterlogged": "false"})
		floor = block.MakeBlock("purpur_slab", map[string]string{"type": "top", "waterlogged": "true"})
	case 1:
		coral = block.MakeBlock("fire_coral_fan", map[string]string{"waterlogged": "false"})
		floor = block.MakeBlock("red_nether_brick_slab", map[string]string{"type": "top", "waterlogged": "true"})
	case 2:
		coral = block.MakeBlock("dead_brain_coral_fan", map[string]string{"waterlogged": "false"})
		floor = block.MakeBlock("gray_wool", nil)
	}

	xdim := 2 + rand.Intn(3)
	zdim := 2 + rand.Intn(3)
	for x := 0; x < xdim; x++ {
		for z := 0; z < zdim; z++ {
			dec.SetBlock(x, 0, z, coral)
			dec.SetBlock(x, -1, z, floor)
		}
	}

	return dec
}

// Table
func CenterDecoration2(meta DecorationMeta) Decoration {
	dec := make(Decoration)

	full_carpet := rand.Intn(5) == 0

	piston_head := block.MakeBlock("piston_head", map[string]string{"facing": "up"})
	dye_color := block.RandMat(block.DYE_COLORS)
	carpet := block.MakeBlock(dye_color+"_carpet", nil)

	makeBorder(&dec, -1, -1, 4, 4)

	for x := 0; x < 4; x++ {
		for z := 0; z < 4; z++ {
			dec.SetBlock(x, 0, z, piston_head)
			dec.SetBlock(x, 1, z, carpet)
		}
	}

	if full_carpet {
		dec.SetBlock(1, 0, 1, block.BARRIER)
		dec.SetBlock(1, 0, 2, block.BARRIER)
		dec.SetBlock(2, 0, 1, block.BARRIER)
		dec.SetBlock(2, 0, 2, block.BARRIER)
	} else {
		var center_block_west, center_block_east block.Block
		switch rand.Intn(4) {
		case 0:
			center_block_west = block.RandBlock(block.LEAF_BLOCKS)
			center_block_east = center_block_west
		case 1:
			center_block_west = block.MakeBlock(dye_color+"_wool", nil)
			center_block_east = center_block_west
		case 2:
			center_block_west = block.MakeBlock(block.RandMat(block.SLAB_MATERIALS)+"_slab", map[string]string{"type": "top"})
			center_block_east = center_block_west
		case 3:
			center_block_west = block.MakeBlock(block.RandMat(block.WOOD_MATERIALS)+"_trapdoor", map[string]string{"facing": "east", "half": "top"})
			center_block_east = block.MakeBlock(block.RandMat(block.WOOD_MATERIALS)+"_trapdoor", map[string]string{"facing": "west", "half": "top"})
		}
		dec.SetBlock(1, 0, 1, center_block_west)
		dec.SetBlock(1, 0, 2, center_block_west)
		dec.SetBlock(2, 0, 1, center_block_east)
		dec.SetBlock(2, 0, 2, center_block_east)
		for x := 1; x <= 2; x++ {
			for z := 1; z <= 2; z++ {
				dec.SetBlock(x, 1, z, GetOneBlockDecoration(0.5))
			}
		}
	}

	if rand.Intn(2) == 0 {
		chair := RandomChair(DecorationMeta{ZDim: 1, IncludeArmrests: true})
		lineChairs(&dec, &chair, -1, 0, 1, 0, 2, 2)
		lineChairs(&dec, &chair, 2, 0, -1, 1, 2, 2)
		lineChairs(&dec, &chair, 4, 0, 2, 2, 2, 2)
		lineChairs(&dec, &chair, 1, 0, 4, 3, 2, 2)
	} else {
		chair := RandomChair(DecorationMeta{ZDim: 2, IncludeArmrests: true})
		dec.ApplyDecoration(-1, 0, 1, chair)
		dec.ApplyDecoration(2, 0, -1, chair.Rotate(1))
		dec.ApplyDecoration(4, 0, 2, chair.Rotate(2))
		dec.ApplyDecoration(1, 0, 4, chair.Rotate(3))
	}

	return dec
}

// Table
func CenterDecoration3(meta DecorationMeta) Decoration {
	dec := make(Decoration)

	wood_material := block.RandMat(block.WOOD_MATERIALS)
	dye_color := block.RandMat(block.DYE_COLORS)

	makeBorder(&dec, -1, -1, 2, 2)

	mat_type := rand.Intn(4)

	for i := 0; i < 4; i++ {
		dir1 := i
		dir2 := (i + 1) % 4
		trapdoorDir := direction.DirectionNames[i]
		fenceDir1 := direction.DirectionNames[(i+2)%4]
		fenceDir2 := direction.DirectionNames[(i+3)%4]
		offset := direction.DirectionOffsets[dir1].Add(direction.DirectionOffsets[dir2])
		x := (offset.X + 1) / 2
		z := (offset.Z + 1) / 2

		var base, top block.Block
		switch mat_type {
		case 0:
			base = block.MakeBlock(wood_material+"_fence", map[string]string{fenceDir1: "true", fenceDir2: "true"})
			top = block.MakeBlock(wood_material+"_pressure_plate", nil)
		case 1:
			base = block.MakeBlock(wood_material+"_fence", map[string]string{fenceDir1: "true", fenceDir2: "true"})
			top = block.MakeBlock(wood_material+"_trapdoor", map[string]string{"facing": trapdoorDir})
		case 2:
			base = block.MakeBlock("iron_bars", map[string]string{fenceDir1: "true", fenceDir2: "true"})
			top = block.MakeBlock("iron_trapdoor", map[string]string{"facing": fenceDir1})
		case 3:
			base = block.MakeBlock(dye_color+"_stained_glass_pane", map[string]string{fenceDir1: "true", fenceDir2: "true"})
			top = block.MakeBlock(dye_color+"_carpet", nil)
		}
		dec.SetBlock(x, 0, z, base)
		dec.SetBlock(x, 1, z, top)
	}

	if rand.Intn(2) == 0 {
		chair := RandomChair(DecorationMeta{ZDim: 1, IncludeArmrests: true})
		lineChairs(&dec, &chair, -1, 0, 0, 0, 2, 2)
		lineChairs(&dec, &chair, 2, 0, 1, 2, 2, 2)
	} else {
		chair := RandomChair(DecorationMeta{ZDim: 2, IncludeArmrests: true})
		dec.ApplyDecoration(-1, 0, 0, chair)
		dec.ApplyDecoration(2, 0, 1, chair.Rotate(2))
	}

	return dec
}

var centerDecorationGenerators = []DecorationGenerator{
	CenterDecoration0,
	CenterDecoration1,
	CenterDecoration2,
	CenterDecoration3,
}

func RandomCenterDecoration(meta DecorationMeta) Decoration {
	generator := centerDecorationGenerators[rand.Intn(len(centerDecorationGenerators))]
	return generator(meta).Rotate(rand.Intn(4))
}

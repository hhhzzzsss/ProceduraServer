package decorations

import (
	"github.com/hhhzzzsss/procedura-generator/structuregen/block"
	"github.com/hhhzzzsss/procedura-generator/util"
)

func DoubleDoors(meta *DecorationMeta) Decoration {
	var dec Decoration
	door_material := randMat(WOOD_DOOR_MATERIALS)
	llDoor := block.MakeBlock(door_material, map[string]string{"facing": "north", "half": "lower", "hinge": "left"})
	ulDoor := block.MakeBlock(door_material, map[string]string{"facing": "north", "half": "upper", "hinge": "left"})
	lrDoor := block.MakeBlock(door_material, map[string]string{"facing": "north", "half": "lower", "hinge": "right"})
	urDoor := block.MakeBlock(door_material, map[string]string{"facing": "north", "half": "upper", "hinge": "right"})
	dec[util.MakeVec3i(0, 0, 0)] = llDoor
	dec[util.MakeVec3i(0, 1, 0)] = ulDoor
	dec[util.MakeVec3i(1, 0, 0)] = lrDoor
	dec[util.MakeVec3i(1, 1, 0)] = urDoor
	for x := 0; x < 2; x++ {
		for y := 0; y < 2; y++ {
			dec[util.MakeVec3i(x, y, -1)] = block.AIR
			dec[util.MakeVec3i(x, y, 1)] = block.AIR
		}
	}
	return dec
}

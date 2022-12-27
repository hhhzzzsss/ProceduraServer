package decorations

import (
	"github.com/hhhzzzsss/procedura-generator/structuregen/block"
	"github.com/hhhzzzsss/procedura-generator/util"
)

func DoubleDoors(meta *DecorationMeta) Decoration {
	dec := make(Decoration)
	door_material := randMat(WOOD_DOOR_MATERIALS)
	llDoor := block.MakeBlock(door_material, map[string]string{"facing": "west", "half": "lower", "hinge": "left"})
	ulDoor := block.MakeBlock(door_material, map[string]string{"facing": "west", "half": "upper", "hinge": "left"})
	lrDoor := block.MakeBlock(door_material, map[string]string{"facing": "west", "half": "lower", "hinge": "right"})
	urDoor := block.MakeBlock(door_material, map[string]string{"facing": "west", "half": "upper", "hinge": "right"})
	dec[util.MakeVec3i(0, 0, 1)] = llDoor
	dec[util.MakeVec3i(0, 1, 1)] = ulDoor
	dec[util.MakeVec3i(0, 0, 0)] = lrDoor
	dec[util.MakeVec3i(0, 1, 0)] = urDoor
	for z := 0; z < 2; z++ {
		for y := 0; y < 2; y++ {
			dec[util.MakeVec3i(-1, y, z)] = block.AIR
			dec[util.MakeVec3i(1, y, z)] = block.AIR
		}
	}
	return dec
}

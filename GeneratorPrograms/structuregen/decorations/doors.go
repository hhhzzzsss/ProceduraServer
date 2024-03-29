package decorations

import (
	"github.com/hhhzzzsss/procedura-generator/structuregen/block"
)

func DoubleDoors(meta DecorationMeta) Decoration {
	dec := make(Decoration)
	door_material := block.RandMat(block.WOOD_MATERIALS) + "_door"
	llDoor := block.MakeBlock(door_material, map[string]string{"facing": "west", "half": "lower", "hinge": "left"})
	ulDoor := block.MakeBlock(door_material, map[string]string{"facing": "west", "half": "upper", "hinge": "left"})
	lrDoor := block.MakeBlock(door_material, map[string]string{"facing": "west", "half": "lower", "hinge": "right"})
	urDoor := block.MakeBlock(door_material, map[string]string{"facing": "west", "half": "upper", "hinge": "right"})
	dec.SetBlock(0, 0, 1, llDoor)
	dec.SetBlock(0, 1, 1, ulDoor)
	dec.SetBlock(0, 0, 0, lrDoor)
	dec.SetBlock(0, 1, 0, urDoor)
	for z := 0; z < 2; z++ {
		for y := 0; y < 2; y++ {
			dec.SetBlock(-1, y, z, block.AIR)
			dec.SetBlock(1, y, z, block.AIR)
		}
	}
	return dec
}

func SingleDoor(meta DecorationMeta) Decoration {
	dec := make(Decoration)
	door_material := block.RandMat(block.WOOD_MATERIALS) + "_door"
	lDoor := block.MakeBlock(door_material, map[string]string{"facing": "west", "half": "lower", "hinge": "right"})
	uDoor := block.MakeBlock(door_material, map[string]string{"facing": "west", "half": "upper", "hinge": "right"})
	dec.SetBlock(0, 0, 0, lDoor)
	dec.SetBlock(0, 1, 0, uDoor)
	for y := 0; y < 2; y++ {
		dec.SetBlock(-1, y, 0, block.AIR)
		dec.SetBlock(1, y, 0, block.AIR)
	}
	return dec
}

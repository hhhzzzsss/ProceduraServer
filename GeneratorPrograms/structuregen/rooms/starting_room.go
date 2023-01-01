package rooms

import (
	"github.com/hhhzzzsss/procedura-generator/structuregen/block"
	"github.com/hhhzzzsss/procedura-generator/structuregen/decorations"
	"github.com/hhhzzzsss/procedura-generator/structuregen/direction"
	"github.com/hhhzzzsss/procedura-generator/util"
)

type StartingRoom struct {
	RoomBase
	roomSize, roomHeight int
}

func (r *StartingRoom) GetRoomBase() *RoomBase {
	return &r.RoomBase
}

func (r *StartingRoom) Initialize(meta RoomMeta) {
	r.ClearAll()

	title1, title1_width := decorations.GetBlockText("mansion", "blackstone", "blackstone_slab", "blackstone_stairs")
	title2, title2_width := decorations.GetBlockText("of babel", "blackstone", "blackstone_slab", "blackstone_stairs")
	credits1, credits1_width := decorations.GetBlockText("made by", "blackstone", "blackstone_slab", "blackstone_stairs")
	credits2, credits2_width := decorations.GetBlockText("hhhzzzsss", "prismarine_bricks", "prismarine_brick_slab", "prismarine_brick_stairs")

	r.roomSize = util.Max(title1_width+2, title2_width+2, credits1_width+4, credits2_width+4)
	r.roomHeight = 12

	r.FillBlocks(0, 0, 0, r.roomSize-1, r.roomHeight-1, r.roomSize-1, block.AIR)
	r.MakeHollowCuboid(0, 0, 0, r.roomSize-1, r.roomHeight-1, r.roomSize-1, block.SMOOTH_QUARTZ)

	for y := 1; y < r.roomHeight-1; y++ {
		for x := 2; x < r.roomSize-2; x++ {
			r.SetReplaceableBlock(x, y, 0, true)
			r.SetReplaceableBlock(x, y, r.roomSize-1, true)
		}
	}

	title1_pos := (r.roomSize - title1_width) / 2
	title2_pos := (r.roomSize - title2_width) / 2
	credits1_pos := (r.roomSize - credits1_width) / 2
	credits2_pos := (r.roomSize - credits2_width) / 2
	r.ApplyDecoration(-1, 8, title1_pos, title1)
	r.ApplyDecoration(-1, 4, title2_pos, title2)
	r.ApplyDecoration(r.roomSize-2, 7, credits1_pos, credits1)
	r.ApplyDecoration(r.roomSize-2, 3, credits2_pos, credits2)

	r.MainEntranceLocation = util.MakeVec3i(0, 1, (r.roomSize-1)/2)
	r.MainEntrance = decorations.DoubleDoors(decorations.DefaultDecorationMeta)
	r.ApplyMainEntrance()

	for i := 2; i < r.roomSize-2; i++ {
		r.AddEntranceLocation(
			i, 1, 0,
			direction.South,
			nil,
			DefaultRoomMeta,
		)
		r.AddEntranceLocation(
			i, 1, r.roomSize-1,
			direction.North,
			nil,
			DefaultRoomMeta,
		)
	}
}

func (r *StartingRoom) Finalize(rv RegionView, meta RoomMeta) {
	for x := 1; x < r.roomSize-1; x++ {
		for z := 1; z < r.roomSize-1; z++ {
			rv.SetBlock(x, r.roomHeight-2, z, block.LIGHT_15)
		}
	}
	for y := 1; y < r.roomHeight-1; y++ {
		for z := 1; z < r.roomSize; z++ {
			x := r.roomSize - 2
			b, ok := r.GetBlock(x, y, z)
			if ok && b.IsAir() {
				rv.SetBlock(x, y, z, block.LIGHT_15)
			}
		}
	}
}

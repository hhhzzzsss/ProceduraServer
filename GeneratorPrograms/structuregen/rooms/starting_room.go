package rooms

import (
	"github.com/hhhzzzsss/procedura-generator/structuregen/block"
	"github.com/hhhzzzsss/procedura-generator/structuregen/decorations"
	"github.com/hhhzzzsss/procedura-generator/structuregen/direction"
	"github.com/hhhzzzsss/procedura-generator/util"
)

type StartingRoom struct {
	RoomBase
}

func (r *StartingRoom) GetRoomBase() *RoomBase {
	return &r.RoomBase
}

func (r *StartingRoom) Initialize(meta *RoomMeta) {
	r.ClearAll()

	title1, title1_width := decorations.GetBlockText("mansion", "blackstone", "blackstone_slab", "blackstone_stairs")
	title2, title2_width := decorations.GetBlockText("of babel", "blackstone", "blackstone_slab", "blackstone_stairs")
	credits1, credits1_width := decorations.GetBlockText("made by", "blackstone", "blackstone_slab", "blackstone_stairs")
	credits2, credits2_width := decorations.GetBlockText("hhhzzzsss", "prismarine_bricks", "prismarine_brick_slab", "prismarine_brick_stairs")

	room_size := util.Max(title1_width+2, title2_width+2, credits1_width+4, credits2_width+4)
	room_height := 12

	r.FillBlocks(0, 0, 0, room_size-1, room_height-1, room_size-1, block.AIR)
	r.MakeHollowCuboid(0, 0, 0, room_size-1, room_height-1, room_size-1, block.SMOOTH_QUARTZ)

	for y := 1; y < room_height-1; y++ {
		for x := 2; x < room_size-2; x++ {
			r.SetReplaceableBlock(x, y, 0, true)
			r.SetReplaceableBlock(x, y, room_size-1, true)
		}
	}

	title1_pos := (room_size - title1_width) / 2
	title2_pos := (room_size - title2_width) / 2
	credits1_pos := (room_size - credits1_width) / 2
	credits2_pos := (room_size - credits2_width) / 2
	r.ApplyDecoration(-1, 8, title1_pos, title1)
	r.ApplyDecoration(-1, 4, title2_pos, title2)
	r.ApplyDecoration(room_size-2, 7, credits1_pos, credits1)
	r.ApplyDecoration(room_size-2, 3, credits2_pos, credits2)

	r.MainEntranceLocation = util.MakeVec3i(0, 1, (room_size-1)/2)
	r.MainEntrance = decorations.DoubleDoors(&decorations.DefaultDecorationMeta)
	r.ApplyMainEntrance()

	for i := 2; i < room_size-2; i++ {
		r.AddEntranceLocation(
			i, 1, 0,
			direction.South,
			nil,
			&DefaultRoomMeta,
		)
		r.AddEntranceLocation(
			i, 1, room_size-1,
			direction.North,
			nil,
			&DefaultRoomMeta,
		)
	}
}

func (r *StartingRoom) Finalize(meta *RoomMeta) {
}

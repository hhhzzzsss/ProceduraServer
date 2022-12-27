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

	const room_size = 16
	const room_height = 8

	// r.AddBoundingBox(0, 0, 0, room_size-1, room_height-1, room_size-1)

	r.FillBlocks(0, 0, 0, room_size-1, room_height-1, room_size-1, block.AIR)
	r.MakeHollowCuboid(0, 0, 0, room_size-1, room_height-1, room_size-1, block.SMOOTH_QUARTZ)

	for y := 1; y < room_height-1; y++ {
		for x := 2; x < room_size-2; x++ {
			r.SetReplaceableBlock(x, y, 0, true)
			r.SetReplaceableBlock(x, y, room_size-1, true)
		}
	}

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

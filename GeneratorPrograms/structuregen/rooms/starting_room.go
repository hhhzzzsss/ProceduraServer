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

	r.AddBoundingBox(0, 0, 0, room_size-1, room_height-1, room_size-1)

	r.MakeHollowCuboid(0, 0, 0, room_size-1, room_height-1, room_size-1, block.MakeBlock("smooth_quartz", nil))

	for y := 1; y < room_height-1; y++ {
		for z := 2; z < room_size-2; z++ {
			r.SetReplacableBlock(0, y, z, true)
			r.SetReplacableBlock(room_size-1, y, z, true)
		}
	}

	r.MainEntranceLocation = util.MakeVec3i((room_size-1)/2, 1, 0)
	r.MainEntrance = decorations.DoubleDoors(&decorations.DefaultDecorationMeta)
	r.ApplyMainEntrance()

	for i := 2; i < room_size-2; i++ {
		r.AddEntranceLocation(
			0, 1, i,
			direction.West,
			nil,
			&DefaultRoomMeta,
		)
		r.AddEntranceLocation(
			room_size-1, 1, i,
			direction.East,
			nil,
			&DefaultRoomMeta,
		)
	}
}

func (r *StartingRoom) Finalize(meta *RoomMeta) {
}

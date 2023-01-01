package rooms

import (
	"github.com/hhhzzzsss/procedura-generator/structuregen/block"
	"github.com/hhhzzzsss/procedura-generator/structuregen/decorations"
	"github.com/hhhzzzsss/procedura-generator/structuregen/direction"
	"github.com/hhhzzzsss/procedura-generator/util"
)

type DefaultRoom struct {
	RoomBase
	roomSize, roomHeight int
}

func (r *DefaultRoom) GetRoomBase() *RoomBase {
	return &r.RoomBase
}

func (r *DefaultRoom) Initialize(meta RoomMeta) {
	r.ClearAll()

	r.roomSize = 16
	r.roomHeight = 8

	r.FillBlocks(0, 0, 0, r.roomSize-1, r.roomHeight-1, r.roomSize-1, block.AIR)
	r.MakeHollowCuboid(0, 0, 0, r.roomSize-1, r.roomHeight-1, r.roomSize-1, block.SMOOTH_QUARTZ)

	for y := 1; y < r.roomHeight-1; y++ {
		for i := 2; i < r.roomSize-2; i++ {
			r.SetReplaceableBlock(i, y, 0, true)
			r.SetReplaceableBlock(i, y, r.roomSize-1, true)
			r.SetReplaceableBlock(r.roomSize-1, y, i, true)
		}
	}

	r.MainEntranceLocation = util.MakeVec3i(-1, 1, (r.roomSize-1)/2)
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
		r.AddEntranceLocation(
			r.roomSize-1, 1, i,
			direction.West,
			nil,
			DefaultRoomMeta,
		)
	}
}

func (r *DefaultRoom) Finalize(rv RegionView, meta RoomMeta) {
	ceilingMeta := decorations.DecorationMeta{XDim: r.roomSize - 4, ZDim: r.roomSize - 4}
	ceilingLight := decorations.RandomCeilingLight(ceilingMeta)
	rv.ApplyDecoration(2, r.roomHeight-2, 2, ceilingLight)
}

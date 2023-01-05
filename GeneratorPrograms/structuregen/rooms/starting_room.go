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
	WallOrigins          [4]util.Vec3i
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
	r.WallOrigins = [4]util.Vec3i{
		util.MakeVec3i(0, 1, 1),
		util.MakeVec3i(r.roomSize-2, 1, 0),
		util.MakeVec3i(r.roomSize-1, 1, r.roomSize-2),
		util.MakeVec3i(1, 1, r.roomSize-1),
	}

	r.FillBlocks(0, 0, 0, r.roomSize-1, r.roomHeight-1, r.roomSize-1, block.AIR)
	r.MakeHollowCuboid(0, 0, 0, r.roomSize-1, r.roomHeight-1, r.roomSize-1, block.SMOOTH_QUARTZ)

	for _, dir := range []int{1, 3} {
		orig := r.WallOrigins[dir]
		offset := direction.DirectionOffsets[(dir+3)%4]
		for i := 1; i < r.roomSize-3; i++ {
			pos := orig.Add(offset.Scale(i))
			for y := 1; y < r.roomHeight-1; y++ {
				r.SetReplaceableBlock(pos.X, y, pos.Z, true)
			}
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

	for _, dir := range []int{1, 3} {
		orig := r.WallOrigins[dir]
		offset := direction.DirectionOffsets[(dir+3)%4]
		entranceDirection := direction.Direction((dir + 2) % 4)
		for i := 1; i < r.roomSize-3; i++ {
			pos := orig.Add(offset.Scale(i))
			r.AddEntranceLocation(
				pos.X, pos.Y, pos.Z,
				entranceDirection,
				nil,
				RoomMeta{LeftWallSpace: i + 1, RightWallSpace: (r.roomSize - 1) - (i + 1)},
			)
		}
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

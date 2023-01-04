package rooms

import (
	"math/rand"

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
	r.FillBlocks(0, -1, 0, r.roomSize-1, -1, r.roomSize-1, block.SMOOTH_QUARTZ)

	for y := 1; y < r.roomHeight-1; y++ {
		for i := 2; i < r.roomSize-2; i++ {
			r.SetReplaceableBlock(i, y, 0, true)
			r.SetReplaceableBlock(i, y, r.roomSize-1, true)
			r.SetReplaceableBlock(r.roomSize-1, y, i, true)
		}
	}
	for x := 1; x < r.roomSize-1; x++ {
		for z := 1; z < r.roomSize-1; z++ {
			r.SetReplaceableBlock(x, 0, z, true)
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

	const max_center_attempts = 50
	const max_center_success = 5
	numCenterDecorationSamples := 2 + rand.Intn(4)

	centerDecorationSamples := make([]decorations.Decoration, 0, numCenterDecorationSamples)
	centerDecorationBoundingBoxes := make([]util.BoundingBox, 0, numCenterDecorationSamples)
	for i := 0; i < numCenterDecorationSamples; i++ {
		dec := decorations.RandomCenterDecoration(decorations.DefaultDecorationMeta)
		centerDecorationSamples = append(centerDecorationSamples, dec)
		centerDecorationBoundingBoxes = append(centerDecorationBoundingBoxes, dec.GetBoundingBox())
	}

	centerPositions := make([]util.Vec3i, 0, (r.roomSize-4)*(r.roomSize-4))
	for x := 2; x < r.roomSize-2; x++ {
		for z := 2; z < r.roomSize-2; z++ {
			centerPositions = append(centerPositions, util.MakeVec3i(x, 1, z))
		}
	}

	for attempt, success := 0, 0; attempt < max_center_attempts && success < max_center_success; attempt++ {
		pos := util.RemoveRandomFromSlice(&centerPositions)
		for i, dec := range centerDecorationSamples {
			bb := centerDecorationBoundingBoxes[i]
			if bb.X1+pos.X < 2 || bb.X2+pos.X >= r.roomSize-2 || bb.Z1+pos.Z < 2 || bb.Z2+pos.Z >= r.roomSize-2 {
				continue
			}
			if !rv.CanPlaceDecoration(pos.X, pos.Y, pos.Z, dec) {
				continue
			}
			rv.ApplyDecoration(pos.X, pos.Y, pos.Z, dec)
			util.RemoveFromUnorderedSlice(&centerDecorationSamples, i)
			util.RemoveFromUnorderedSlice(&centerDecorationBoundingBoxes, i)
			success++
			break
		}
	}
}

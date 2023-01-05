package rooms

import (
	"math/rand"

	"github.com/hhhzzzsss/procedura-generator/structuregen/block"
	"github.com/hhhzzzsss/procedura-generator/structuregen/decorations"
	"github.com/hhhzzzsss/procedura-generator/structuregen/direction"
	"github.com/hhhzzzsss/procedura-generator/util"
)

type Stairs struct {
	RoomBase
	attachedRoom *RoomView
	attachedMeta RoomMeta
}

func (r *Stairs) GetRoomBase() *RoomBase {
	return &r.RoomBase
}

func (r *Stairs) Initialize(meta RoomMeta) {
	if meta.LeftWallSpace < 1 || meta.RightWallSpace < 2 {
		return
	}

	r.ClearAll()

	heightChange := 8 + rand.Intn(8)
	ascending := rand.Intn(2) == 0
	var increment int
	if ascending {
		increment = 1
	} else {
		increment = -1
	}

	fullBlock := block.SMOOTH_QUARTZ
	var stair block.Block
	if ascending {
		stair = block.MakeBlock("smooth_quartz_stairs", map[string]string{"facing": "east"})
	} else {
		stair = block.MakeBlock("smooth_quartz_stairs", map[string]string{"facing": "west"})
	}
	lantern := block.MakeBlock("lantern", map[string]string{"hanging": "true"})

	for z := 0; z < 4; z++ {
		r.SetBlock(0, -1, z, fullBlock)
		r.SetBlock(0, 4, z, fullBlock)
	}
	for y := 0; y < 4; y++ {
		r.SetBlock(0, y, 0, fullBlock)
		r.SetBlock(0, y, 3, fullBlock)
	}
	for z := 1; z < 3; z++ {
		for y := 0; y < 4; y++ {
			r.SetBlock(0, y, z, block.AIR)
		}
	}

	for i := 0; i < heightChange; i++ {
		basex := i + 1
		basey := i * increment
		if !ascending {
			basey--
		}
		for z := 0; z < 4; z++ {
			r.SetBlock(basex, -1+basey, z, fullBlock)
			r.SetBlock(basex, 5+basey, z, fullBlock)
		}
		for y := 0; y < 5; y++ {
			r.SetBlock(basex, basey+y, 0, fullBlock)
			r.SetBlock(basex, basey+y, 3, fullBlock)
		}
		r.SetBlock(basex, basey, 1, stair)
		r.SetBlock(basex, basey, 2, stair)

		for z := 1; z < 3; z++ {
			for y := 1; y < 5; y++ {
				r.SetBlock(basex, basey+y, z, block.AIR)
			}
		}
		if i%3 == 0 {
			if i%6 == 0 {
				r.SetBlock(basex, basey+4, 1, lantern)
			} else {
				r.SetBlock(basex, basey+4, 2, lantern)
			}
		}
	}

	for z := 0; z < 4; z++ {
		r.SetBlock(heightChange+1, -1+heightChange*increment, z, fullBlock)
		r.SetBlock(heightChange+1, 4+heightChange*increment, z, fullBlock)
	}
	for y := 0; y < 4; y++ {
		r.SetBlock(heightChange+1, heightChange*increment+y, 0, fullBlock)
		r.SetBlock(heightChange+1, heightChange*increment+y, 3, fullBlock)
	}
	for z := 1; z < 3; z++ {
		for y := 0; y < 4; y++ {
			r.SetBlock(heightChange+1, heightChange*increment+y, z, block.AIR)
		}
	}

	for z := 0; z < 4; z++ {
		for y := -1; y < 5; y++ {
			r.SetBlock(heightChange+2, heightChange*increment+y, z, fullBlock)
		}
	}
	for z := 1; z < 3; z++ {
		for y := 0; y < 4; y++ {
			r.SetReplaceableBlock(heightChange+2, y+heightChange*increment, z, true)
		}
	}

	r.MainEntranceLocation = util.MakeVec3i(-1, 0, 1)
	r.MainEntrance = decorations.DoubleDoors(decorations.DefaultDecorationMeta)
	r.ApplyMainEntrance()

	pseudoEntranceLocation := util.MakeVec3i(heightChange+2, heightChange*increment, 1)
	attachedRoomGenerator := func() ([]Room, []float32) {
		rooms := []Room{
			&DefaultRoom{},
		}
		weights := []float32{
			1.0,
		}
		return rooms, weights
	}
	r.attachedMeta = meta
	r.attachedMeta.LeftWallSpace = 1
	r.attachedMeta.RightWallSpace = 2
	r.attachedMeta.Elevation = meta.Elevation + heightChange*increment
	r.attachedRoom = r.attachRoom(pseudoEntranceLocation, direction.West, attachedRoomGenerator, r.attachedMeta)

	if r.attachedRoom == nil {
		r.Invalid = true
	}
}

func (r *Stairs) attachRoom(pseudoEntranceLocation util.Vec3i, entranceDir direction.Direction, attachedRoomGenerator func() ([]Room, []float32), attachedMeta RoomMeta) *RoomView {
	possibleRooms, possibleRoomWeights := attachedRoomGenerator()

generateRoomOuterLoop:
	for attempts := 0; attempts < 2 && len(possibleRooms) > 0; attempts++ {
		room := util.RemoveWeightedRandomFromSlice(&possibleRooms, &possibleRoomWeights)

		room.Initialize(attachedMeta)
		if room.GetRoomBase().Invalid {
			continue
		}

		roomView := GetRoomView(room, pseudoEntranceLocation, entranceDir)
		transformedMainEntrance := roomView.GetTransformedMainEntranceExterior()
		for pos := range transformedMainEntrance {
			if !r.CanReplaceBlock(pos) {
				continue generateRoomOuterLoop
			}
		}
		transformedPositions, transformedBlocks := roomView.GetTransformedBlocks()
		transformedRPositions, transformedRBlocks := roomView.GetTransformedReplaceableBlocks()
		for _, pos := range transformedPositions {
			r, r_ok := r.ReplaceableBlocks[pos]
			if r_ok && !r {
				continue generateRoomOuterLoop
			}
		}

		for pos, block := range transformedMainEntrance {
			r.Blocks[pos] = block
			r.ReplaceableBlocks[pos] = false
		}
		for i, pos := range transformedPositions {
			r.Blocks[pos] = transformedBlocks[i]
		}
		for i, pos := range transformedRPositions {
			r.ReplaceableBlocks[pos] = transformedRBlocks[i]
		}

		r.EntranceLocations = roomView.GetEntranceLocations()

		return &roomView
	}
	return nil
}

func (r *Stairs) Finalize(rv RegionView, meta RoomMeta) {
	for pos, block := range r.Blocks {
		attachedRoomPos := r.attachedRoom.InvTransformVec(pos)
		r.attachedRoom.Room.GetRoomBase().Blocks[attachedRoomPos] = block
	}
	for pos, state := range r.ReplaceableBlocks {
		attachedRoomPos := r.attachedRoom.InvTransformVec(pos)
		r.attachedRoom.Room.GetRoomBase().ReplaceableBlocks[attachedRoomPos] = state
	}
	newAttachedRoomView := GetRoomView(r.attachedRoom.Room, rv.RoomView.TransformVec(r.attachedRoom.Pos), direction.Direction(r.attachedRoom.Dir).Rotate(rv.RoomView.Dir))
	attachedRegionView := GetRegionView(rv.Region, &newAttachedRoomView)
	r.attachedRoom.Room.Finalize(attachedRegionView, r.attachedMeta)
}

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

	if meta.AboveGround {
		r.generateWindows(rv)
	}
}

func (r *DefaultRoom) generateWindows(rv RegionView) {
	var wallOrigins = [4]util.Vec3i{
		util.MakeVec3i(0, 1, 1),
		util.MakeVec3i(r.roomSize-2, 1, 0),
		util.MakeVec3i(r.roomSize-1, 1, r.roomSize-2),
		util.MakeVec3i(1, 1, r.roomSize-1),
	}

	wallSections := make([]WallSection, 0)

	for dir := 0; dir < 4; dir++ {
		orig := wallOrigins[dir]
		perpendicularOffset := direction.DirectionOffsets[dir]
		parallelOffset := direction.DirectionOffsets[(dir+3)%4]
		sectionStart := 0
	wallLoop:
		for i := 0; i < r.roomSize-2; i++ {
			columnBase := orig.Add(parallelOffset.Scale(i))
			for j := 0; j < r.roomHeight-2; j++ {
				wallPos := columnBase.Add(util.MakeVec3i(0, j, 0))
				outsidePos := wallPos.Add(perpendicularOffset)
				if !r.CanReplaceBlock(wallPos) || (j > 0 && j < r.roomHeight-3 && !rv.IsEmpty(outsidePos.X, outsidePos.Y, outsidePos.Z)) {
					sectionWidth := i - sectionStart
					if sectionWidth > 0 {
						sectionOrigin := orig.Add(parallelOffset.Scale(sectionStart))
						wallSections = append(wallSections, MakeWallSection(sectionOrigin.X, sectionOrigin.Y, sectionOrigin.Z, sectionWidth, r.roomHeight-2, dir))
					}
					sectionStart = i + 1
					continue wallLoop
				}
			}
		}
		sectionWidth := (r.roomSize - 2) - sectionStart
		if sectionWidth > 0 {
			sectionOrigin := orig.Add(parallelOffset.Scale(sectionStart))
			wallSections = append(wallSections, MakeWallSection(sectionOrigin.X, sectionOrigin.Y, sectionOrigin.Z, sectionWidth, r.roomHeight-2, dir))
		}
	}

	for _, section := range wallSections {
		if section.Width < 4 {
			continue
		}
		if rand.Intn(2) == 0 {
			rv.FillWallSection(section, 0, 0, 1, 1, block.GLASS)
			if rand.Intn(2) == 0 {
				windowsill := decorations.Windowsill(decorations.DecorationMeta{ZDim: section.Width}).Rotate(section.Dir)
				pos := util.MakeVec3i(section.X, section.Y, section.Z).Sub(direction.DirectionOffsets[section.Dir])
				if rv.CanPlaceDecoration(pos.X, pos.Y, pos.Z, windowsill) {
					rv.ApplyDecoration(pos.X, pos.Y, pos.Z, windowsill)
				}
			}
		}
	}
}

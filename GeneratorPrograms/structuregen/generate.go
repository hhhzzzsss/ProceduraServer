package structuregen

import (
	"math/rand"

	"github.com/hhhzzzsss/procedura-generator/region"
	"github.com/hhhzzzsss/procedura-generator/structuregen/direction"
	"github.com/hhhzzzsss/procedura-generator/structuregen/rooms"
	"github.com/hhhzzzsss/procedura-generator/util"
)

type StructureGenSettings struct {
	XDim, YDim, ZDim          int
	XOrigin, YOrigin, ZOrigin int
	StartingEntranceDirection direction.Direction
	StartingRoomGenerator     func() []rooms.Room
	StartingRoomMeta          rooms.RoomMeta
	MaxRoomAttempts           int
}

func defaultRoomGenerator() []rooms.Room {
	return []rooms.Room{
		&rooms.DefaultRoom{},
	}
}

func GetDefaultSettings() StructureGenSettings {
	return StructureGenSettings{
		XDim: 256, YDim: 256, ZDim: 256,
		XOrigin: 127, YOrigin: 1, ZOrigin: 0,
		StartingEntranceDirection: direction.North,
		StartingRoomGenerator:     defaultRoomGenerator,
		MaxRoomAttempts:           5,
	}
}

func GenerateStructure(settings *StructureGenSettings) region.Region {
	xdim, ydim, zdim := settings.XDim, settings.YDim, settings.ZDim
	region := region.MakeRegion(xdim, ydim, zdim)
	region.AddPaletteBlock("air")

	origin_entrance := &rooms.EntranceLocation{
		Pos:           util.MakeVec3i(settings.XOrigin, settings.YOrigin, settings.ZOrigin),
		Dir:           settings.StartingEntranceDirection,
		RoomGenerator: settings.StartingRoomGenerator,
		Meta:          &settings.StartingRoomMeta,
	}
	potentialEntrances := []*rooms.EntranceLocation{origin_entrance}

	roomViews := make([]*rooms.RoomView, 0)

	for len(potentialEntrances) > 0 {
		// Select random entrance location and remove from potentialEntrances
		selectedEntranceLocation := removeRandomFromSlice(potentialEntrances)

		// Generate room and add new RoomView / EntranceLocation(s) to lists if exists
		rv := generateRoom(selectedEntranceLocation, &region, settings)
		if rv != nil {
			roomViews = append(roomViews, rv)
			for _, el := range rv.GetEntranceLocations() {
				potentialEntrances = append(potentialEntrances, &el)
			}
		}
	}

	return region
}

func generateRoom(entranceLocation *rooms.EntranceLocation, region *region.Region, settings *StructureGenSettings) *rooms.RoomView {
	possibleRooms := entranceLocation.RoomGenerator()
	parent := entranceLocation.Parent
generateRoomOuterLoop:
	for attempts := 0; attempts < settings.MaxRoomAttempts && len(possibleRooms) > 0; attempts++ {
		room := removeRandomFromSlice(possibleRooms)
		room.Initialize(entranceLocation.Meta)
		roomView := rooms.GetView(room, entranceLocation.Pos, entranceLocation.Dir)
		transformedMainEntrance := roomView.GetTransformedMainEntranceExterior()
		if entranceLocation.Parent != nil {
			for pos := range transformedMainEntrance {
				if !parent.Contains(pos) || !parent.CanReplaceBlock(pos) {
					continue generateRoomOuterLoop
				}
			}
			for pos, block := range transformedMainEntrance {
				parent.ReplaceBlock(pos, block)
			}
		}
		transformedPositions, transformedBlocks := roomView.GetTransformedBlocks()
		for _, pos := range transformedPositions {
			regionVal := region.Get(pos.X, pos.Y, pos.Z)
			parentVal, ok := parent.GetBlock(pos)
			if regionVal != 0 && !(ok && parentVal.IsAir()) {
				continue generateRoomOuterLoop
			}
		}
		for i, pos := range transformedPositions {
			region.SetWithName(pos.X, pos.Y, pos.Z, transformedBlocks[i].ToString())
		}
	}
	return nil
}

func removeRandomFromSlice[T any](s []T) T {
	selectedIdx := rand.Intn(len(s))
	lastIdx := len(s) - 1
	selectedElem := s[selectedIdx]
	s[selectedIdx] = s[lastIdx]
	s = s[:lastIdx]
	return selectedElem
}

package structuregen

import (
	"fmt"
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
	StartingRoomGenerator     func() ([]rooms.Room, []float32)
	StartingRoomMeta          rooms.RoomMeta
	MaxRoomAttempts           int
}

func defaultStartingRoomGenerator() ([]rooms.Room, []float32) {
	rooms := []rooms.Room{
		&rooms.StartingRoom{},
	}
	weights := []float32{
		1.0,
	}
	return rooms, weights
}

func defaultRoomGenerator() ([]rooms.Room, []float32) {
	rooms := []rooms.Room{
		&rooms.DefaultRoom{},
	}
	weights := []float32{
		1.0,
	}
	return rooms, weights
}

func GetDefaultSettings() StructureGenSettings {
	return StructureGenSettings{
		XDim: 256, YDim: 256, ZDim: 256,
		XOrigin: 1, YOrigin: 50, ZOrigin: 127,
		StartingEntranceDirection: direction.West,
		StartingRoomGenerator:     defaultStartingRoomGenerator,
		MaxRoomAttempts:           5,
	}
}

func GenerateStructure(settings *StructureGenSettings) region.Region {
	xdim, ydim, zdim := settings.XDim, settings.YDim, settings.ZDim

	// Create palettte for the superflat terrain and fill region with structure void
	fmt.Println("Initializing region...")
	region := region.MakeRegion(xdim, ydim, zdim)
	region.AddPaletteBlock("structure_void") // 0
	region.AddPaletteBlock("air")            // 1
	region.AddPaletteBlock("grass_block")    // 2
	region.AddPaletteBlock("dirt")           // 3
	region.AddPaletteBlock("stone")          // 4
	region.AddPaletteBlock("bedrock")        // 5
	region.ForEach(func(x, y, z int) int {
		return 0
	})

	fmt.Println("Generating Structure...")

	origin_entrance := rooms.EntranceLocation{
		Pos:           util.MakeVec3i(settings.XOrigin, settings.YOrigin, settings.ZOrigin),
		Dir:           settings.StartingEntranceDirection,
		RoomGenerator: settings.StartingRoomGenerator,
		Meta:          &settings.StartingRoomMeta,
	}
	potentialEntrances := []rooms.EntranceLocation{origin_entrance}

	roomViews := make([]*rooms.RoomView, 0)

	for len(potentialEntrances) > 0 {
		// Select random entrance location and remove from potentialEntrances
		selectedEntranceLocation := removeRandomFromSlice(&potentialEntrances)

		// Generate room and add new RoomView / EntranceLocation(s) to lists if exists
		rv := generateRoom(selectedEntranceLocation, &region, settings)
		if rv != nil {
			roomViews = append(roomViews, rv)
			potentialEntrances = append(potentialEntrances, rv.GetEntranceLocations()...)
		}
	}

	fmt.Println("Filling superflat terrain...")
	region.ForEach(func(x, y, z int) int {
		if region.Get(x, y, z) == 0 {
			if y >= 50 {
				return 1 // air
			} else if y == 49 {
				return 2 // grass_block
			} else if y >= 17 && y <= 48 {
				return 3 // dirt
			} else if y >= 1 && y <= 16 {
				return 4 // stone
			} else {
				return 5 // bedrock
			}
		} else {
			return region.Get(x, y, z)
		}
	})

	return region
}

func generateRoom(entranceLocation rooms.EntranceLocation, region *region.Region, settings *StructureGenSettings) *rooms.RoomView {
	var possibleRooms []rooms.Room
	var possibleRoomWeights []float32
	if entranceLocation.RoomGenerator == nil {
		possibleRooms, possibleRoomWeights = defaultRoomGenerator()
	} else {
		possibleRooms, possibleRoomWeights = entranceLocation.RoomGenerator()
	}
	parent := entranceLocation.Parent
	xdim, ydim, zdim := settings.XDim, settings.YDim, settings.ZDim
generateRoomOuterLoop:
	for attempts := 0; attempts < settings.MaxRoomAttempts && len(possibleRooms) > 0; attempts++ {
		room := removeWeightedRandomFromSlice(&possibleRooms, &possibleRoomWeights)
		room.Initialize(entranceLocation.Meta)
		roomView := rooms.GetView(room, entranceLocation.Pos, entranceLocation.Dir)
		transformedMainEntrance := roomView.GetTransformedMainEntranceExterior()
		if parent != nil {
			for pos := range transformedMainEntrance {
				if !parent.CanReplaceBlock(pos) {
					continue generateRoomOuterLoop
				}
			}
		}
		transformedPositions, transformedBlocks := roomView.GetTransformedBlocks()
		for _, pos := range transformedPositions {
			if pos.X < 0 || pos.Y < 0 || pos.Z < 0 || pos.X >= xdim || pos.Y >= ydim || pos.Z >= zdim {
				continue generateRoomOuterLoop
			}
			regionVal := region.Get(pos.X, pos.Y, pos.Z)
			if regionVal != 0 {
				continue generateRoomOuterLoop
			}
		}
		if parent != nil {
			for pos, block := range transformedMainEntrance {
				parent.ReplaceBlock(pos, block)
			}
		}
		for pos, block := range transformedMainEntrance {
			// Needs extra check for the nil parent case
			if pos.X >= 0 && pos.Y >= 0 && pos.Z >= 0 && pos.X < xdim && pos.Y < ydim && pos.Z < zdim {
				region.SetWithName(pos.X, pos.Y, pos.Z, block.ToString())
			}
		}
		for i, pos := range transformedPositions {
			region.SetWithName(pos.X, pos.Y, pos.Z, transformedBlocks[i].ToString())
		}
		return &roomView
	}
	return nil
}

func removeRandomFromSlice[T any](s *[]T) T {
	selectedIdx := rand.Intn(len(*s))
	lastIdx := len(*s) - 1
	selectedElem := (*s)[selectedIdx]
	(*s)[selectedIdx] = (*s)[lastIdx]
	(*s) = (*s)[:lastIdx]
	return selectedElem
}

func removeWeightedRandomFromSlice[T any](s *[]T, weights *[]float32) T {
	if len(*s) != len(*weights) {
		panic("Element slice must have same length as weight slice")
	}

	lastIdx := len(*s) - 1

	var totalWeight float32 = 0
	for _, weight := range *weights {
		totalWeight += weight
	}

	rval := rand.Float32() * totalWeight
	var cumWeight float32 = 0
	for i, weight := range *weights {
		cumWeight += weight
		if cumWeight >= rval {
			selectedElem := (*s)[i]
			(*s)[i] = (*s)[lastIdx]
			(*s) = (*s)[:lastIdx]
			(*weights)[i] = (*weights)[lastIdx]
			(*weights) = (*weights)[:lastIdx]
			return selectedElem
		}
	}
	selectedElem := (*s)[lastIdx]
	(*s) = (*s)[:lastIdx]
	(*weights) = (*weights)[:lastIdx]
	return selectedElem
}

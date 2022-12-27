package rooms

import (
	"github.com/hhhzzzsss/procedura-generator/structuregen/block"
	"github.com/hhhzzzsss/procedura-generator/structuregen/decorations"
	"github.com/hhhzzzsss/procedura-generator/structuregen/direction"
	"github.com/hhhzzzsss/procedura-generator/util"
)

/*
By default, rooms will have the entrance at the origin, and the main entrance will
point toward the negative x direction (West). The structure gen code will handle
rotations with RoomView, but the room struct only sees itself in the default
orientation.
*/

type Room interface {
	// All rooms must have a room base
	GetRoomBase() *RoomBase

	/*
		Constructs everything in the room, including:
		- Bounding boxes
		- Blocks
		- Replaceable blocks
		- Main entrance / main entrance location
		- Possible entrance locations connecting to other rooms
		- All decorations that apply prior to adding other adjacent rooms
	*/
	Initialize(meta *RoomMeta)

	// Whatever decorations/blocks to add after adding other adjacent rooms
	Finalize(meta *RoomMeta)
}

type RoomBase struct {
	// BoundingBoxes        []util.BoundingBox
	Blocks               map[util.Vec3i]block.Block
	ReplaceableBlocks    map[util.Vec3i]bool // 1 is initial state, set to 0 when replaced
	MainEntrance         decorations.Decoration
	MainEntranceLocation util.Vec3i
	EntranceLocations    []EntranceLocation
}

type RoomMeta struct {
	SolidFacingEntrance bool
}

var DefaultRoomMeta RoomMeta = RoomMeta{
	SolidFacingEntrance: true,
}

type EntranceLocation struct {
	Pos           util.Vec3i          // Base position of the entrance
	Dir           direction.Direction // Direction of the entrance (relative to the potential room that will be put here)
	RoomGenerator func() []Room       // Returns a slice of possible rooms to put here
	Meta          *RoomMeta           // Additional info for room generation
	Parent        *RoomView           // RoomView that is offering the EntranceLocation
}

type RoomView struct {
	Room Room
	Pos  util.Vec3i
	Dir  int
}

func GetView(r Room, origin util.Vec3i, dir direction.Direction) RoomView {
	return RoomView{r, origin, int(dir)}
}

// func (r *RoomBase) CreateBasicBoundingBox() {
// 	bb := util.BoundingBox{
// 		X1: math.MaxInt, Y1: math.MaxInt, Z1: math.MaxInt,
// 		X2: math.MinInt, Y2: math.MinInt, Z2: math.MinInt,
// 	}
// 	for pos := range r.Blocks {
// 		if pos.X < bb.X1 {
// 			bb.X1 = pos.X
// 		}
// 		if pos.Y < bb.Y1 {
// 			bb.Y1 = pos.Y
// 		}
// 		if pos.Z < bb.Z1 {
// 			bb.Z1 = pos.Z
// 		}
// 		if pos.X > bb.X2 {
// 			bb.X2 = pos.X
// 		}
// 		if pos.Y > bb.Y2 {
// 			bb.Y2 = pos.Y
// 		}
// 		if pos.Z > bb.Z2 {
// 			bb.Z2 = pos.Z
// 		}
// 	}

// 	if bb.X2 < bb.X1 || bb.Y2 < bb.Y1 || bb.Z2 < bb.Z1 {
// 		panic("Tried to create basic bounding box, but room has no blocks")
// 	}

// 	r.BoundingBoxes = []util.BoundingBox{bb}
// }

// func (r *RoomBase) AddBoundingBox(x1, y1, z1, x2, y2, z2 int) {
// 	bb := util.BoundingBox{
// 		X1: x1, Y1: y1, Z1: z1,
// 		X2: x2, Y2: y2, Z2: z2,
// 	}
// 	r.BoundingBoxes = append(r.BoundingBoxes, bb)
// }

// Sets the block at a position
func (r *RoomBase) SetBlock(x, y, z int, block block.Block) {
	r.Blocks[util.MakeVec3i(x, y, z)] = block
}

// Fills a volume with a particular block.
// Upper and lower bounds are inclusive.
func (r *RoomBase) FillBlocks(x1, y1, z1, x2, y2, z2 int, block block.Block) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			for z := z1; z <= z2; z++ {
				r.Blocks[util.MakeVec3i(x, y, z)] = block
			}
		}
	}
}

// Creates a hollow cuboid with the specified bounds.
// Upper and lower bounds are inclusive.
func (r *RoomBase) MakeHollowCuboid(x1, y1, z1, x2, y2, z2 int, block block.Block) {
	r.FillBlocks(x1, y1, z1, x2, y2, z1, block)
	r.FillBlocks(x1, y1, z2, x2, y2, z2, block)
	r.FillBlocks(x1, y1, z1, x2, y1, z2, block)
	r.FillBlocks(x1, y2, z1, x2, y2, z2, block)
	r.FillBlocks(x1, y1, z1, x1, y2, z2, block)
	r.FillBlocks(x2, y1, z1, x2, y2, z2, block)
}

// Adds a replaceable block flag at the specified position
func (r *RoomBase) SetReplaceableBlock(x, y, z int, b bool) {
	r.ReplaceableBlocks[util.MakeVec3i(x, y, z)] = b
}

// Fills a volume with a replaceable blockf lag
// Upper and lower bounds are inclusive.
func (r *RoomBase) FillReplaceableBlocks(x1, y1, z1, x2, y2, z2 int, b bool) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			for z := z1; z <= z2; z++ {
				r.ReplaceableBlocks[util.MakeVec3i(x, y, z)] = b
			}
		}
	}
}

func (r *RoomBase) ApplyMainEntrance() {
	for pos, block := range r.MainEntrance {
		roomPos := r.MainEntranceLocation.Add(pos)
		if r.Contains(roomPos) {
			r.Blocks[roomPos] = block
			r.ReplaceableBlocks[roomPos] = false
		}
	}
}

func (r *RoomBase) AddEntranceLocation(
	X, Y, Z int,
	Dir direction.Direction,
	GetPossibleRooms func() []Room,
	Meta *RoomMeta,
) {
	r.EntranceLocations = append(r.EntranceLocations, EntranceLocation{util.MakeVec3i(X, Y, Z), Dir, GetPossibleRooms, Meta, nil})
}

// func (r *RoomBase) ClearBoundingBoxes() {
// 	r.BoundingBoxes = make([]util.BoundingBox, 0)
// }

func (r *RoomBase) ClearBlocks() {
	r.Blocks = make(map[util.Vec3i]block.Block)
}

func (r *RoomBase) ClearReplaceableBlocks() {
	r.ReplaceableBlocks = make(map[util.Vec3i]bool)
}

func (r *RoomBase) ClearEntranceLocations() {
	r.EntranceLocations = make([]EntranceLocation, 0)
}

func (r *RoomBase) ClearAll() {
	// r.ClearBoundingBoxes()
	r.ClearBlocks()
	r.ClearReplaceableBlocks()
	r.ClearEntranceLocations()
}

func (r *RoomBase) Contains(pos util.Vec3i) bool {
	// for _, bb := range r.BoundingBoxes {
	// 	if bb.Contains(pos) {
	// 		return true
	// 	}
	// }
	// return false
	_, ok := r.Blocks[pos]
	return ok
}

// Transforms a vector from room space to global space
func (rv *RoomView) TransformVec(v util.Vec3i) util.Vec3i {
	v = v.Sub(rv.Room.GetRoomBase().MainEntranceLocation)
	v = direction.RotateVec(v, rv.Dir)
	v = v.Add(rv.Pos)
	return v
}

// Transforms a vector from global space to room space
func (rv *RoomView) InvTransformVec(v util.Vec3i) util.Vec3i {
	v = v.Sub(rv.Pos)
	v = direction.RotateVec(v, -rv.Dir)
	v = v.Add(rv.Room.GetRoomBase().MainEntranceLocation)
	return v
}

func (rv *RoomView) GetBlock(v util.Vec3i) (block.Block, bool) {
	block, ok := rv.Room.GetRoomBase().Blocks[rv.InvTransformVec(v)]
	return block, ok
}

func (rv *RoomView) GetReplaceableBlock(v util.Vec3i) (bool, bool) {
	block, ok := rv.Room.GetRoomBase().ReplaceableBlocks[rv.InvTransformVec(v)]
	return block, ok
}

func (rv *RoomView) CanReplaceBlock(v util.Vec3i) bool {
	vTrans := rv.InvTransformVec(v)
	b := rv.Room.GetRoomBase().Blocks[vTrans]
	r, r_ok := rv.Room.GetRoomBase().ReplaceableBlocks[vTrans]

	if r_ok {
		return r
	} else {
		return b.IsAir()
	}
}

func (rv *RoomView) ReplaceBlock(v util.Vec3i, b block.Block) {
	vTrans := rv.InvTransformVec(v)
	rv.Room.GetRoomBase().Blocks[vTrans] = b
	rv.Room.GetRoomBase().ReplaceableBlocks[vTrans] = false
}

func (rv *RoomView) GetEntranceLocations() []EntranceLocation {
	transformedEntranceLocations := make([]EntranceLocation, len(rv.Room.GetRoomBase().EntranceLocations))
	for i, el := range rv.Room.GetRoomBase().EntranceLocations {
		transformedEntranceLocations[i] = EntranceLocation{
			rv.TransformVec(el.Pos),
			el.Dir.Rotate(int(rv.Dir)),
			el.RoomGenerator,
			el.Meta,
			rv,
		}
	}
	return transformedEntranceLocations
}

func (rv *RoomView) GetTransformedMainEntranceExterior() decorations.Decoration {
	newDecoration := make(decorations.Decoration)
	roomBase := rv.Room.GetRoomBase()
	for pos, block := range roomBase.MainEntrance {
		roomPos := rv.Room.GetRoomBase().MainEntranceLocation.Add(pos)
		if !roomBase.Contains(roomPos) {
			newDecoration[rv.TransformVec(roomPos)] = block.Rotate(rv.Dir)
		}
	}
	return newDecoration
}

// Returns a slice of block positions and a slice of the corresponding blocks
func (rv *RoomView) GetTransformedBlocks() ([]util.Vec3i, []block.Block) {
	positions := make([]util.Vec3i, 0)
	blocks := make([]block.Block, 0)
	for pos, block := range rv.Room.GetRoomBase().Blocks {
		positions = append(positions, rv.TransformVec(pos))
		blocks = append(blocks, block.Rotate(rv.Dir))
	}
	return positions, blocks
}

func (rv *RoomView) Contains(v util.Vec3i) bool {
	return rv.Room.GetRoomBase().Contains(rv.InvTransformVec(v))
}

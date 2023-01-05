package rooms

import (
	"math/rand"

	"github.com/hhhzzzsss/procedura-generator/structuregen/block"
	"github.com/hhhzzzsss/procedura-generator/structuregen/decorations"
	"github.com/hhhzzzsss/procedura-generator/util"
)

type Balcony struct {
	RoomBase
	width, depth int
}

func (r *Balcony) GetRoomBase() *RoomBase {
	return &r.RoomBase
}

func (r *Balcony) Initialize(meta RoomMeta) {
	if meta.LeftWallSpace < 3 || meta.RightWallSpace < 3 || rand.Intn(10) > 0 {
		return
	}

	r.ClearAll()

	r.width = 7
	r.depth = 4

	wood_material := block.RandMat(block.WOOD_MATERIALS)
	platform_block := block.MakeBlock(wood_material+"_planks", nil)
	x_axis_fence := block.MakeBlock(wood_material+"_fence", map[string]string{"west": "true", "east": "true"})
	z_axis_fence := block.MakeBlock(wood_material+"_fence", map[string]string{"north": "true", "south": "true"})
	corner_fence_1 := block.MakeBlock(wood_material+"_fence", map[string]string{"west": "true", "south": "true"})
	corner_fence_2 := block.MakeBlock(wood_material+"_fence", map[string]string{"west": "true", "north": "true"})
	lantern := block.MakeBlock("lantern", map[string]string{"hanging": "false"})

	r.FillBlocks(0, 0, 0, r.depth-1, 5, r.width-1, block.AIR)
	r.FillBlocks(0, 0, 0, r.depth-1, 0, r.width-1, platform_block)

	for x := 0; x < r.depth-1; x++ {
		r.SetBlock(x, 1, 0, x_axis_fence)
		r.SetBlock(x, 1, r.width-1, x_axis_fence)
	}
	for z := 1; z < r.width-1; z++ {
		r.SetBlock(r.depth-1, 1, z, z_axis_fence)
	}
	r.SetBlock(r.depth-1, 1, 0, corner_fence_1)
	r.SetBlock(r.depth-1, 1, r.width-1, corner_fence_2)
	r.SetBlock(r.depth-1, 2, 0, lantern)
	r.SetBlock(r.depth-1, 2, r.width-1, lantern)

	r.MainEntranceLocation = util.MakeVec3i(-1, 1, (r.width-1)/2)
	r.MainEntrance = decorations.SingleDoor(decorations.DefaultDecorationMeta)
	r.ApplyMainEntrance()
}

func (r *Balcony) Finalize(rv RegionView, meta RoomMeta) {
}

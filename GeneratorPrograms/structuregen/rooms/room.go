package rooms

import (
	"github.com/hhhzzzsss/procedura-generator/structuregen"
	"github.com/hhhzzzsss/procedura-generator/structuregen/decorations"
	"github.com/hhhzzzsss/procedura-generator/util"
)

type Room interface {
	GetEntrances() []decorations.Decoration
}

type RoomBase struct {
	BoundingBoxes    []util.BoundingBox
	Origin           util.Vec3i
	Blocks           map[util.Vec3i]structuregen.Block
	ReplacableBlocks map[util.Vec3i]bool // 1 is initial state, set to 0 when replaced
}

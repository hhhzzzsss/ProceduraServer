package decorations

import (
	"github.com/hhhzzzsss/procedura-generator/structuregen/block"
	"github.com/hhhzzzsss/procedura-generator/structuregen/direction"
	"github.com/hhhzzzsss/procedura-generator/util"
)

type Decoration map[util.Vec3i]block.Block
type DecorationMeta struct {
	xdim, ydim, zdim int
}

var DefaultDecorationMeta DecorationMeta = DecorationMeta{}

type DecorationGenerator func(meta *DecorationMeta) Decoration

func (d Decoration) Rotate(a int) Decoration {
	newDecoration := make(Decoration)
	for pos, block := range d {
		newDecoration[direction.RotateVec(pos, a)] = block.Rotate(a)
	}
	return newDecoration
}

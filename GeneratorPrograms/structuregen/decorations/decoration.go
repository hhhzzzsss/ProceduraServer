package decorations

import (
	"github.com/hhhzzzsss/procedura-generator/structuregen"
	"github.com/hhhzzzsss/procedura-generator/util"
)

type Decoration map[util.Vec3i]structuregen.Block
type DecorationMeta struct {
	xdim, ydim, zdim int
}

type DecorationGenerator func(meta *DecorationMeta) Decoration

func (d *Decoration) Rotate(a int) {
	newDecoration := make(Decoration)
	for pos, block := range *d {
		newDecoration[structuregen.RotateVec(pos, a)] = block.Rotate(a)
	}
	*d = newDecoration
}

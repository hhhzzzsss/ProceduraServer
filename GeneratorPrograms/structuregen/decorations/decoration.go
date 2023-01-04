package decorations

import (
	"math"

	"github.com/hhhzzzsss/procedura-generator/structuregen/block"
	"github.com/hhhzzzsss/procedura-generator/structuregen/direction"
	"github.com/hhhzzzsss/procedura-generator/util"
)

type Decoration map[util.Vec3i]block.Block
type DecorationMeta struct {
	XDim, YDim, ZDim int
	IncludeArmrests  bool
}

var DefaultDecorationMeta DecorationMeta = DecorationMeta{}

type DecorationGenerator func(meta DecorationMeta) Decoration

func (d Decoration) Rotate(a int) Decoration {
	newDecoration := make(Decoration)
	for pos, block := range d {
		newDecoration[direction.RotateVec(pos, a)] = block.Rotate(a)
	}
	return newDecoration
}

func (d Decoration) GetBoundingBox() util.BoundingBox {
	bb := util.BoundingBox{
		X1: math.MaxInt, Y1: math.MaxInt, Z1: math.MaxInt,
		X2: math.MinInt, Y2: math.MinInt, Z2: math.MinInt,
	}

	for pos, block := range d {
		if block.IsAir() {
			continue
		}
		if pos.X < bb.X1 {
			bb.X1 = pos.X
		}
		if pos.Y < bb.Y1 {
			bb.Y1 = pos.Y
		}
		if pos.Z < bb.Z1 {
			bb.Z1 = pos.Z
		}
		if pos.X > bb.X2 {
			bb.X2 = pos.X
		}
		if pos.Y > bb.Y2 {
			bb.Y2 = pos.Y
		}
		if pos.Z > bb.Z2 {
			bb.Z2 = pos.Z
		}
	}

	if bb.X2 < bb.X1 {
		panic("Cannot get bounding box for empty decoration")
	}

	return bb
}

func (d *Decoration) SetBlock(x, y, z int, block block.Block) {
	(*d)[util.MakeVec3i(x, y, z)] = block
}

func (d *Decoration) ApplyDecoration(x, y, z int, dec Decoration) {
	centerPos := util.MakeVec3i(x, y, z)
	for pos, block := range dec {
		(*d)[centerPos.Add(pos)] = block
	}
}

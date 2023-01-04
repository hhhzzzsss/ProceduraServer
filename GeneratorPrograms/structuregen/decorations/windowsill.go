package decorations

import "github.com/hhhzzzsss/procedura-generator/structuregen/block"

func Windowsill(meta DecorationMeta) Decoration {
	dec := make(Decoration)

	sill := block.MakeBlock("smooth_quartz_stairs", map[string]string{"facing": "west", "half": "top"})
	for z := 0; z < meta.ZDim; z++ {
		dec.SetBlock(0, 0, z, sill)
		dec.SetBlock(0, 1, z, GetWindowsillBlockDecoration(0.5))
	}

	return dec
}

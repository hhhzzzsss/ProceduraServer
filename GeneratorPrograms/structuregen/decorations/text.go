package decorations

import (
	"strings"

	"github.com/hhhzzzsss/procedura-generator/structuregen/block"
)

// Returns Decoration containing block text and text width
func GetBlockText(text, blockname, slabname, stairname string) (Decoration, int) {
	zpos := 0
	blk := block.MakeBlock(blockname, nil)
	bslab := block.MakeBlock(slabname, map[string]string{"type": "bottom"})
	tslab := block.MakeBlock(slabname, map[string]string{"type": "top"})
	blstair := block.MakeBlock(stairname, map[string]string{"half": "bottom", "facing": "north"})
	brstair := block.MakeBlock(stairname, map[string]string{"half": "bottom", "facing": "south"})
	tlstair := block.MakeBlock(stairname, map[string]string{"half": "top", "facing": "north"})
	trstair := block.MakeBlock(stairname, map[string]string{"half": "top", "facing": "south"})

	dec := make(Decoration)

	for _, c := range strings.ToLower(text) {
		if c == 'a' {
			dec.SetBlock(0, 0, zpos+0, blk)
			dec.SetBlock(0, 1, zpos+0, blk)
			dec.SetBlock(0, 2, zpos+0, brstair)
			dec.SetBlock(0, 1, zpos+1, bslab)
			dec.SetBlock(0, 2, zpos+1, tslab)
			dec.SetBlock(0, 0, zpos+2, blk)
			dec.SetBlock(0, 1, zpos+2, blk)
			dec.SetBlock(0, 2, zpos+2, blstair)
			zpos += 4
		} else if c == 'b' {
			dec.SetBlock(0, 0, zpos+0, blk)
			dec.SetBlock(0, 1, zpos+0, blk)
			dec.SetBlock(0, 2, zpos+0, blk)
			dec.SetBlock(0, 0, zpos+1, bslab)
			dec.SetBlock(0, 1, zpos+1, tslab)
			dec.SetBlock(0, 2, zpos+1, tslab)
			dec.SetBlock(0, 0, zpos+2, tlstair)
			dec.SetBlock(0, 1, zpos+2, blstair)
			dec.SetBlock(0, 2, zpos+2, blstair)
			zpos += 4
		} else if c == 'c' {
			dec.SetBlock(0, 0, zpos+0, trstair)
			dec.SetBlock(0, 1, zpos+0, blk)
			dec.SetBlock(0, 2, zpos+0, brstair)
			dec.SetBlock(0, 0, zpos+1, bslab)
			dec.SetBlock(0, 2, zpos+1, tslab)
			dec.SetBlock(0, 0, zpos+2, tlstair)
			dec.SetBlock(0, 2, zpos+2, blstair)
			zpos += 4
		} else if c == 'd' {
			dec.SetBlock(0, 0, zpos+0, blk)
			dec.SetBlock(0, 1, zpos+0, blk)
			dec.SetBlock(0, 2, zpos+0, blk)
			dec.SetBlock(0, 0, zpos+1, bslab)
			dec.SetBlock(0, 2, zpos+1, tslab)
			dec.SetBlock(0, 0, zpos+2, tlstair)
			dec.SetBlock(0, 1, zpos+2, blk)
			dec.SetBlock(0, 2, zpos+2, blstair)
			zpos += 4
		} else if c == 'e' {
			dec.SetBlock(0, 0, zpos+0, blk)
			dec.SetBlock(0, 1, zpos+0, blk)
			dec.SetBlock(0, 2, zpos+0, blk)
			dec.SetBlock(0, 0, zpos+1, bslab)
			dec.SetBlock(0, 1, zpos+1, bslab)
			dec.SetBlock(0, 2, zpos+1, tslab)
			dec.SetBlock(0, 0, zpos+2, bslab)
			dec.SetBlock(0, 2, zpos+2, tslab)
			zpos += 4
		} else if c == 'f' {
			dec.SetBlock(0, 0, zpos+0, blk)
			dec.SetBlock(0, 1, zpos+0, blk)
			dec.SetBlock(0, 2, zpos+0, blk)
			dec.SetBlock(0, 1, zpos+1, bslab)
			dec.SetBlock(0, 2, zpos+1, tslab)
			dec.SetBlock(0, 2, zpos+2, tslab)
			zpos += 4
		} else if c == 'g' {
			dec.SetBlock(0, 0, zpos+0, trstair)
			dec.SetBlock(0, 1, zpos+0, blk)
			dec.SetBlock(0, 2, zpos+0, brstair)
			dec.SetBlock(0, 0, zpos+1, bslab)
			dec.SetBlock(0, 2, zpos+1, tslab)
			dec.SetBlock(0, 0, zpos+2, brstair)
			dec.SetBlock(0, 1, zpos+2, bslab)
			dec.SetBlock(0, 2, zpos+2, blstair)
			zpos += 4
		} else if c == 'h' {
			dec.SetBlock(0, 0, zpos+0, blk)
			dec.SetBlock(0, 1, zpos+0, blk)
			dec.SetBlock(0, 2, zpos+0, blk)
			dec.SetBlock(0, 1, zpos+1, bslab)
			dec.SetBlock(0, 0, zpos+2, blk)
			dec.SetBlock(0, 1, zpos+2, blk)
			dec.SetBlock(0, 2, zpos+2, blk)
			zpos += 4
		} else if c == 'i' {
			dec.SetBlock(0, 0, zpos+0, bslab)
			dec.SetBlock(0, 2, zpos+0, tslab)
			dec.SetBlock(0, 0, zpos+1, blk)
			dec.SetBlock(0, 1, zpos+1, blk)
			dec.SetBlock(0, 2, zpos+1, blk)
			dec.SetBlock(0, 0, zpos+2, bslab)
			dec.SetBlock(0, 2, zpos+2, tslab)
			zpos += 4
		} else if c == 'j' {
			dec.SetBlock(0, 0, zpos+0, trstair)
			dec.SetBlock(0, 0, zpos+1, bslab)
			dec.SetBlock(0, 0, zpos+2, tlstair)
			dec.SetBlock(0, 1, zpos+2, blk)
			dec.SetBlock(0, 2, zpos+2, blk)
			zpos += 4
		} else if c == 'k' {
			dec.SetBlock(0, 0, zpos+0, blk)
			dec.SetBlock(0, 1, zpos+0, blk)
			dec.SetBlock(0, 2, zpos+0, blk)
			dec.SetBlock(0, 1, zpos+1, blk)
			dec.SetBlock(0, 0, zpos+2, trstair)
			dec.SetBlock(0, 1, zpos+2, tlstair)
			dec.SetBlock(0, 2, zpos+2, brstair)
			zpos += 4
		} else if c == 'l' {
			dec.SetBlock(0, 0, zpos+0, blk)
			dec.SetBlock(0, 1, zpos+0, blk)
			dec.SetBlock(0, 2, zpos+0, blk)
			dec.SetBlock(0, 0, zpos+1, bslab)
			dec.SetBlock(0, 0, zpos+2, bslab)
			zpos += 4
		} else if c == 'm' {
			dec.SetBlock(0, 0, zpos+0, blk)
			dec.SetBlock(0, 1, zpos+0, blk)
			dec.SetBlock(0, 2, zpos+0, blstair)
			dec.SetBlock(0, 1, zpos+1, blk)
			dec.SetBlock(0, 0, zpos+2, blk)
			dec.SetBlock(0, 1, zpos+2, blk)
			dec.SetBlock(0, 2, zpos+2, brstair)
			zpos += 4
		} else if c == 'n' {
			dec.SetBlock(0, 0, zpos+0, blk)
			dec.SetBlock(0, 1, zpos+0, blk)
			dec.SetBlock(0, 2, zpos+0, blstair)
			dec.SetBlock(0, 1, zpos+1, blstair)
			dec.SetBlock(0, 0, zpos+2, trstair)
			dec.SetBlock(0, 1, zpos+2, blk)
			dec.SetBlock(0, 2, zpos+2, blk)
			zpos += 4
		} else if c == 'o' {
			dec.SetBlock(0, 0, zpos+0, trstair)
			dec.SetBlock(0, 1, zpos+0, blk)
			dec.SetBlock(0, 2, zpos+0, brstair)
			dec.SetBlock(0, 0, zpos+1, bslab)
			dec.SetBlock(0, 2, zpos+1, tslab)
			dec.SetBlock(0, 0, zpos+2, tlstair)
			dec.SetBlock(0, 1, zpos+2, blk)
			dec.SetBlock(0, 2, zpos+2, blstair)
			zpos += 4
		} else if c == 'p' {
			dec.SetBlock(0, 0, zpos+0, blk)
			dec.SetBlock(0, 1, zpos+0, blk)
			dec.SetBlock(0, 2, zpos+0, blk)
			dec.SetBlock(0, 1, zpos+1, bslab)
			dec.SetBlock(0, 2, zpos+1, tslab)
			dec.SetBlock(0, 1, zpos+2, tlstair)
			dec.SetBlock(0, 2, zpos+2, blstair)
			zpos += 4
		} else if c == 'q' {
			dec.SetBlock(0, 0, zpos+0, trstair)
			dec.SetBlock(0, 1, zpos+0, blk)
			dec.SetBlock(0, 2, zpos+0, brstair)
			dec.SetBlock(0, 0, zpos+1, brstair)
			dec.SetBlock(0, 2, zpos+1, tslab)
			dec.SetBlock(0, 0, zpos+2, blk)
			dec.SetBlock(0, 1, zpos+2, blk)
			dec.SetBlock(0, 2, zpos+2, blstair)
			zpos += 4
		} else if c == 'r' {
			dec.SetBlock(0, 0, zpos+0, blk)
			dec.SetBlock(0, 1, zpos+0, blk)
			dec.SetBlock(0, 2, zpos+0, blk)
			dec.SetBlock(0, 1, zpos+1, bslab)
			dec.SetBlock(0, 2, zpos+1, tslab)
			dec.SetBlock(0, 0, zpos+2, blstair)
			dec.SetBlock(0, 1, zpos+2, tlstair)
			dec.SetBlock(0, 2, zpos+2, blstair)
			zpos += 4
		} else if c == 's' {
			dec.SetBlock(0, 0, zpos+0, bslab)
			dec.SetBlock(0, 2, zpos+0, brstair)
			dec.SetBlock(0, 0, zpos+1, bslab)
			dec.SetBlock(0, 1, zpos+1, tslab)
			dec.SetBlock(0, 2, zpos+1, tslab)
			dec.SetBlock(0, 0, zpos+2, tlstair)
			dec.SetBlock(0, 1, zpos+2, blstair)
			dec.SetBlock(0, 2, zpos+2, tslab)
			zpos += 4
		} else if c == 't' {
			dec.SetBlock(0, 2, zpos+0, tslab)
			dec.SetBlock(0, 0, zpos+1, blk)
			dec.SetBlock(0, 1, zpos+1, blk)
			dec.SetBlock(0, 2, zpos+1, blk)
			dec.SetBlock(0, 2, zpos+2, tslab)
			zpos += 4
		} else if c == 'u' {
			dec.SetBlock(0, 0, zpos+0, trstair)
			dec.SetBlock(0, 1, zpos+0, blk)
			dec.SetBlock(0, 2, zpos+0, blk)
			dec.SetBlock(0, 0, zpos+1, blk)
			dec.SetBlock(0, 0, zpos+2, tlstair)
			dec.SetBlock(0, 1, zpos+2, blk)
			dec.SetBlock(0, 2, zpos+2, blk)
			zpos += 4
		} else if c == 'v' {
			dec.SetBlock(0, 1, zpos+0, trstair)
			dec.SetBlock(0, 2, zpos+0, blk)
			dec.SetBlock(0, 0, zpos+1, blk)
			dec.SetBlock(0, 1, zpos+1, bslab)
			dec.SetBlock(0, 1, zpos+2, tlstair)
			dec.SetBlock(0, 2, zpos+2, blk)
			zpos += 4
		} else if c == 'w' {
			dec.SetBlock(0, 0, zpos+0, trstair)
			dec.SetBlock(0, 1, zpos+0, blk)
			dec.SetBlock(0, 2, zpos+0, blk)
			dec.SetBlock(0, 1, zpos+1, blk)
			dec.SetBlock(0, 0, zpos+2, tlstair)
			dec.SetBlock(0, 1, zpos+2, blk)
			dec.SetBlock(0, 2, zpos+2, blk)
			zpos += 4
		} else if c == 'x' {
			dec.SetBlock(0, 0, zpos+0, brstair)
			dec.SetBlock(0, 2, zpos+0, trstair)
			dec.SetBlock(0, 0, zpos+1, tslab)
			dec.SetBlock(0, 1, zpos+1, blk)
			dec.SetBlock(0, 2, zpos+1, bslab)
			dec.SetBlock(0, 0, zpos+2, blstair)
			dec.SetBlock(0, 2, zpos+2, tlstair)
			zpos += 4
		} else if c == 'y' {
			dec.SetBlock(0, 2, zpos+0, trstair)
			dec.SetBlock(0, 0, zpos+1, blk)
			dec.SetBlock(0, 1, zpos+1, blk)
			dec.SetBlock(0, 2, zpos+1, bslab)
			dec.SetBlock(0, 2, zpos+2, tlstair)
			zpos += 4
		} else if c == 'z' {
			dec.SetBlock(0, 0, zpos+0, brstair)
			dec.SetBlock(0, 2, zpos+0, tslab)
			dec.SetBlock(0, 0, zpos+1, blk)
			dec.SetBlock(0, 1, zpos+1, brstair)
			dec.SetBlock(0, 2, zpos+1, tslab)
			dec.SetBlock(0, 0, zpos+2, bslab)
			dec.SetBlock(0, 1, zpos+2, tlstair)
			dec.SetBlock(0, 2, zpos+2, blk)
			zpos += 4
		} else {
			zpos += 2
		}
	}

	return dec, zpos - 1
}

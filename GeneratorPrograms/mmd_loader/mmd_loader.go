package mmd_loader

import (
	"bufio"
	"encoding/binary"
	"os"

	"github.com/hhhzzzsss/procedura-generator/region"
	"github.com/hhhzzzsss/procedura-generator/util"
)

const dump_path string = "MMD_DUMP"

func loadDump(cache region.RegionCache[*util.Color]) {
	f, err := os.Open(dump_path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	bufr := bufio.NewReader(f)
	cache.ForEach(func(x, y, z int) *util.Color {
		color := &util.Color{}
		binary.Read(bufr, binary.BigEndian, color)
		return color
	})
}

func LoadDumpAsRegion() region.Region {
	kdtree := util.GetBlockColorKDTree()
	r := region.MakeRegion(256, 256, 256)
	cache := region.MakeRegionCache[*util.Color](&r)
	loadDump(cache)
	r.AddPaletteBlock("air")
	indexMap := make(map[string]int)
	r.ForEach(func(x, y, z int) int {
		if cache.Get(x, y, z).A < 0.0 {
			return 0
		}
		blockColor := kdtree.NearestNeighbor(cache.Get(x, y, z)).(*util.BlockColor)
		blockIdx := r.PaletteSize()
		if idx, ok := indexMap[blockColor.Block]; ok {
			blockIdx = idx
		} else {
			indexMap[blockColor.Block] = blockIdx
			r.AddPaletteBlock(blockColor.Block)
		}
		return blockIdx
	})
	return r
}

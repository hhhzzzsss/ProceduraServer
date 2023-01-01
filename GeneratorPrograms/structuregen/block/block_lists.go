package block

import "math/rand"

var AIR Block = MakeBlock("air", nil)

var SMOOTH_QUARTZ Block = MakeBlock("smooth_quartz", nil)

var GLASS Block = MakeBlock("glass", nil)

var LIGHT_15 = MakeBlock("light", map[string]string{"level": "15"})

var GLOWSTONE = MakeBlock("glowstone", nil)
var SEA_LANTERN = MakeBlock("sea_lantern", nil)
var SHROOMLIGHT = MakeBlock("shroomlight", nil)
var BEACON = MakeBlock("beacon", nil)
var JACK_O_LANTERN = MakeBlock("jack_o_lantern", nil)
var LAVA_CAULDRON = MakeBlock("lava_cauldron", map[string]string{"level": "3"})

var LIGHT_BLOCKS = []Block{
	GLOWSTONE,
	SEA_LANTERN,
	SHROOMLIGHT,
}

var PERMISSIVE_LIGHT_BLOCKS = []Block{
	GLOWSTONE,
	SEA_LANTERN,
	SHROOMLIGHT,
	BEACON,
	JACK_O_LANTERN,
	LAVA_CAULDRON,
}

func RandBlock(list []Block) Block {
	return list[rand.Intn(len(list))]
}

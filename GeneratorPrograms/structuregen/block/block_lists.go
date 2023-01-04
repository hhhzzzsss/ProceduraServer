package block

import "math/rand"

var AIR Block = MakeBlock("air", nil)
var BARRIER Block = MakeBlock("barrier", nil)
var GLASS Block = MakeBlock("glass", nil)
var SMOOTH_QUARTZ Block = MakeBlock("smooth_quartz", nil)

var LIGHT_15 = MakeBlock("light", map[string]string{"level": "15"})
var GLOWSTONE = MakeBlock("glowstone", nil)
var SEA_LANTERN = MakeBlock("sea_lantern", nil)
var SHROOMLIGHT = MakeBlock("shroomlight", nil)
var BEACON = MakeBlock("beacon", nil)
var JACK_O_LANTERN = MakeBlock("jack_o_lantern", nil)
var LAVA_CAULDRON = MakeBlock("lava_cauldron", map[string]string{"level": "3"})

var FLOWER_POT = MakeBlock("flower_pot", nil)
var POTTED_DANDELION = MakeBlock("potted_dandelion", nil)
var POTTED_POPPY = MakeBlock("potted_poppy", nil)
var POTTED_BLUE_ORCHID = MakeBlock("potted_blue_orchid", nil)
var POTTED_ALLIUM = MakeBlock("potted_allium", nil)
var POTTED_AZURE_BLUET = MakeBlock("potted_azure_bluet", nil)
var POTTED_RED_TULIP = MakeBlock("potted_red_tulip", nil)
var POTTED_ORANGE_TULIP = MakeBlock("potted_orange_tulip", nil)
var POTTED_WHITE_TULIP = MakeBlock("potted_white_tulip", nil)
var POTTED_PINK_TULIP = MakeBlock("potted_pink_tulip", nil)
var POTTED_OXEYE_DAISY = MakeBlock("potted_oxeye_daisy", nil)
var POTTED_CORNFLOWER = MakeBlock("potted_cornflower", nil)
var POTTED_LILY_OF_THE_VALLEY = MakeBlock("potted_lily_of_the_valley", nil)
var POTTED_WITHER_ROSE = MakeBlock("potted_wither_rose", nil)
var POTTED_OAK_SAPLING = MakeBlock("potted_oak_sapling", nil)
var POTTED_SPRUCE_SAPLING = MakeBlock("potted_spruce_sapling", nil)
var POTTED_BIRCH_SAPLING = MakeBlock("potted_birch_sapling", nil)
var POTTED_JUNGLE_SAPLING = MakeBlock("potted_jungle_sapling", nil)
var POTTED_ACACIA_SAPLING = MakeBlock("potted_acacia_sapling", nil)
var POTTED_DARK_OAK_SAPLING = MakeBlock("potted_dark_oak_sapling", nil)
var POTTED_RED_MUSHROOM = MakeBlock("potted_red_mushroom", nil)
var POTTED_BROWN_MUSHROOM = MakeBlock("potted_brown_mushroom", nil)
var POTTED_FERN = MakeBlock("potted_fern", nil)
var POTTED_DEAD_BUSH = MakeBlock("potted_dead_bush", nil)
var POTTED_CACTUS = MakeBlock("potted_cactus", nil)
var POTTED_BAMBOO = MakeBlock("potted_bamboo", nil)
var POTTED_AZALEA_BUSH = MakeBlock("potted_azalea_bush", nil)
var POTTED_FLOWERING_AZALEA_BUSH = MakeBlock("potted_flowering_azalea_bush", nil)
var POTTED_CRIMSON_FUNGUS = MakeBlock("potted_crimson_fungus", nil)
var POTTED_WARPED_FUNGUS = MakeBlock("potted_warped_fungus", nil)
var POTTED_CRIMSON_ROOTS = MakeBlock("potted_crimson_roots", nil)
var POTTED_WARPED_ROOTS = MakeBlock("potted_warped_roots", nil)

var OAK_LEAVES = MakeBlock("oak_leaves", map[string]string{"persistent": "true"})
var SPRUCE_LEAVES = MakeBlock("spruce_leaves", map[string]string{"persistent": "true"})
var BIRCH_LEAVES = MakeBlock("birch_leaves", map[string]string{"persistent": "true"})
var JUNGLE_LEAVES = MakeBlock("jungle_leaves", map[string]string{"persistent": "true"})
var ACACIA_LEAVES = MakeBlock("acacia_leaves", map[string]string{"persistent": "true"})
var DARK_OAK_LEAVES = MakeBlock("dark_oak_leaves", map[string]string{"persistent": "true"})
var AZALEA_LEAVES = MakeBlock("azalea_leaves", map[string]string{"persistent": "true"})
var FLOWERING_AZALEA_LEAVES = MakeBlock("flowering_azalea_leaves", map[string]string{"persistent": "true"})

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

var FLOWER_POTS = []Block{
	FLOWER_POT,
	POTTED_DANDELION,
	POTTED_POPPY,
	POTTED_BLUE_ORCHID,
	POTTED_ALLIUM,
	POTTED_AZURE_BLUET,
	POTTED_RED_TULIP,
	POTTED_ORANGE_TULIP,
	POTTED_WHITE_TULIP,
	POTTED_PINK_TULIP,
	POTTED_OXEYE_DAISY,
	POTTED_CORNFLOWER,
	POTTED_LILY_OF_THE_VALLEY,
	POTTED_WITHER_ROSE,
	POTTED_OAK_SAPLING,
	POTTED_SPRUCE_SAPLING,
	POTTED_BIRCH_SAPLING,
	POTTED_JUNGLE_SAPLING,
	POTTED_ACACIA_SAPLING,
	POTTED_DARK_OAK_SAPLING,
	POTTED_RED_MUSHROOM,
	POTTED_BROWN_MUSHROOM,
	POTTED_FERN,
	POTTED_DEAD_BUSH,
	POTTED_CACTUS,
	POTTED_BAMBOO,
	POTTED_AZALEA_BUSH,
	POTTED_FLOWERING_AZALEA_BUSH,
	POTTED_CRIMSON_FUNGUS,
	POTTED_WARPED_FUNGUS,
	POTTED_CRIMSON_ROOTS,
	POTTED_WARPED_ROOTS,
}

var LEAF_BLOCKS = []Block{
	OAK_LEAVES,
	SPRUCE_LEAVES,
	BIRCH_LEAVES,
	JUNGLE_LEAVES,
	ACACIA_LEAVES,
	DARK_OAK_LEAVES,
	AZALEA_LEAVES,
	FLOWERING_AZALEA_LEAVES,
}

func RandBlock(list []Block) Block {
	return list[rand.Intn(len(list))]
}

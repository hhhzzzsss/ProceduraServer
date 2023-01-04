package block

import "math/rand"

var WOOD_MATERIALS = []string{
	"oak",
	"spruce",
	"birch",
	"jungle",
	"acacia",
	"dark_oak",
	"crimson",
	"warped",
}

var SLAB_MATERIALS = []string{
	"oak",
	"spruce",
	"birch",
	"jungle",
	"acacia",
	"dark_oak",
	"crimson",
	"warped",
	"stone",
	"stone_brick",
	"polished_andesite",
	"polished_diorite",
	"polished_granite",
	"smooth_sandstone",
	"smooth_red_sandstone",
	"brick",
	"prismarine",
	"prismarine_brick",
	"dark_prismarine",
	"nether_brick",
	"red_nether_brick",
	"quartz",
	"smooth_quartz",
	"purpur",
	"polished_blackstone",
	"polished_blackstone_brick",
	"waxed_cut_copper",
	"waxed_oxidized_cut_copper",
	"polished_deepslate",
	"deepslate_brick",
	"deepslate_tile",
}

var DYE_COLORS = []string{
	"white",
	"orange",
	"magenta",
	"light_blue",
	"yellow",
	"lime",
	"pink",
	"gray",
	"light_gray",
	"cyan",
	"purple",
	"blue",
	"brown",
	"green",
	"red",
	"black",
}

var GRAYSCALE_COLORS = []string{
	"white",
	"light_gray",
	"gray",
	"black",
}

var CORAL_TYPES = []string{
	"tube",
	"brain",
	"bubble",
	"fire",
	"horn",
}

var MOB_HEADS = []string{
	"skeleton_skull",
	"wither_skeleton_skull",
	"zombie_head",
	"creeper_head",
}

func RandMat(list []string) string {
	return list[rand.Intn(len(list))]
}

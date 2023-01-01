package decorations

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

func RandMat(list []string) string {
	return list[rand.Intn(len(list))]
}

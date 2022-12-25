package decorations

import "math/rand"

var WOOD_DOOR_MATERIALS = []string{
	"oak_door",
	"spruce_door",
	"birch_door",
	"jungle_door",
	"acacia_door",
	"dark_oak_door",
	"crimson_door",
	"warped_door",
}

func randMat(list []string) string {
	return list[rand.Intn(len(list))]
}

package decorations

import (
	"math/rand"
	"strconv"

	"github.com/hhhzzzsss/procedura-generator/structuregen/block"
	"github.com/hhhzzzsss/procedura-generator/structuregen/direction"
)

func GetOneBlockDecoration(air_probability float32) block.Block {
	if rand.Float32() < air_probability {
		return block.AIR
	}

	switch rand.Intn(14) {
	case 0, 1, 2:
		return block.RandBlock(block.FLOWER_POTS)
	case 3, 4:
		candle_material := "candle"
		if rand.Intn(3) > 0 {
			candle_material = block.RandMat(block.DYE_COLORS) + "_candle"
		}
		candle_dict := make(map[string]string)
		candle_dict["candles"] = strconv.Itoa(1 + rand.Intn(4))
		if rand.Intn(4) == 0 {
			candle_dict["lit"] = "false"
		} else {
			candle_dict["lit"] = "true"
		}
		return block.MakeBlock(candle_material, candle_dict)
	case 5:
		return block.MakeBlock("turtle_egg", map[string]string{"eggs": strconv.Itoa(rand.Intn(4) + 1)})
	case 6:
		return block.MakeBlock("sea_pickle", map[string]string{"pickles": strconv.Itoa(rand.Intn(4) + 1), "waterlogged": "false"})
	case 7:
		if rand.Intn(2) == 0 {
			return block.MakeBlock("lantern", nil)
		} else {
			return block.MakeBlock("soul_lantern", nil)
		}
	case 8:
		if rand.Intn(2) == 0 {
			return block.MakeBlock("torch", nil)
		} else {
			return block.MakeBlock("soul_torch", nil)
		}
	case 9:
		return block.MakeBlock("end_rod", map[string]string{"facing": "up"})
	case 10:
		return block.MakeBlock("amethyst_cluster", nil)
	case 11:
		coral_type := block.RandMat(block.CORAL_TYPES)
		state_dict := map[string]string{"waterlogged": "false"}
		if rand.Intn(2) == 0 {
			return block.MakeBlock("dead_"+coral_type+"_coral", state_dict)
		} else {
			return block.MakeBlock("dead_"+coral_type+"_coral_fan", state_dict)
		}
	case 12:
		head_type := block.RandMat(block.MOB_HEADS)
		directionVal := strconv.Itoa(rand.Intn(16))
		state_dict := map[string]string{"rotation": directionVal}
		return block.MakeBlock(head_type, state_dict)
	case 13:
		facingVal := direction.DirectionNames[rand.Intn(4)]
		state_dict := map[string]string{"facing": facingVal}
		switch rand.Intn(3) {
		case 0:
			state_dict["face"] = "floor"
			return block.MakeBlock("grindstone", state_dict)
		case 1:
			return block.MakeBlock("stonecutter", state_dict)
		case 2:
			return block.MakeBlock("bell", state_dict)
		}
	}

	return block.AIR
}

func GetWindowsillBlockDecoration(air_probability float32) block.Block {
	if rand.Float32() < air_probability {
		return block.AIR
	}

	switch rand.Intn(14) {
	case 0, 1, 2, 4, 5:
		return block.RandBlock(block.FLOWER_POTS)
	case 6, 7:
		candle_material := "candle"
		if rand.Intn(3) > 0 {
			candle_material = block.RandMat(block.DYE_COLORS) + "_candle"
		}
		candle_dict := make(map[string]string)
		candle_dict["candles"] = strconv.Itoa(1 + rand.Intn(4))
		if rand.Intn(4) == 0 {
			candle_dict["lit"] = "false"
		} else {
			candle_dict["lit"] = "true"
		}
		return block.MakeBlock(candle_material, candle_dict)
	case 8:
		return block.MakeBlock("turtle_egg", map[string]string{"eggs": strconv.Itoa(rand.Intn(4) + 1)})
	case 9:
		return block.MakeBlock("sea_pickle", map[string]string{"pickles": strconv.Itoa(rand.Intn(4) + 1), "waterlogged": "false"})
	case 10:
		if rand.Intn(2) == 0 {
			return block.MakeBlock("lantern", nil)
		} else {
			return block.MakeBlock("soul_lantern", nil)
		}
	case 11:
		if rand.Intn(2) == 0 {
			return block.MakeBlock("torch", nil)
		} else {
			return block.MakeBlock("soul_torch", nil)
		}
	case 12:
		return block.MakeBlock("end_rod", map[string]string{"facing": "up"})
	case 13:
		return block.MakeBlock("amethyst_cluster", nil)
	case 14:
		coral_type := block.RandMat(block.CORAL_TYPES)
		state_dict := map[string]string{"waterlogged": "false"}
		if rand.Intn(2) == 0 {
			return block.MakeBlock("dead_"+coral_type+"_coral", state_dict)
		} else {
			return block.MakeBlock("dead_"+coral_type+"_coral_fan", state_dict)
		}
	case 15:
		head_type := block.RandMat(block.MOB_HEADS)
		directionVal := strconv.Itoa(rand.Intn(16))
		state_dict := map[string]string{"rotation": directionVal}
		return block.MakeBlock(head_type, state_dict)
	}

	return block.AIR
}

package block

import (
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/hhhzzzsss/procedura-generator/structuregen/direction"
)

type Block struct {
	name  string
	state map[string]string
}

var blockCache map[string]Block

func MakeBlock(name string, state map[string]string) Block {
	newBlock := Block{
		name,
		state,
	}
	if len(newBlock.state) == 0 {
		newBlock.state = nil
	}
	cachedBlock, ok := blockCache[newBlock.ToString()]
	if ok {
		return cachedBlock
	} else {
		return newBlock
	}
}

func (b Block) Rotate(a int) Block {
	if a != 0 && len(b.state) != 0 {
		newState := make(map[string]string)
		for key, value := range b.state {
			switch key {
			case "facing":
				newState["facing"] = direction.RotateDirectionString(value, a)
			case "axis":
				if a == 1 || a == 3 {
					if value == "z" {
						newState["axis"] = "x"
					} else if value == "x" {
						newState["axis"] = "z"
					} else {
						newState["axis"] = value
					}
				} else {
					newState["axis"] = value
				}
			case "rotation":
				dir, err := strconv.Atoi(value)
				if err != nil {
					panic(err.Error())
				}
				newState["rotation"] = strconv.Itoa((dir + 4*a) % 16)
			default:
				newState[direction.RotateDirectionString(key, a)] = value
			}
		}
		return MakeBlock(b.name, newState)
	} else {
		return b
	}
}

func (b Block) ToString() string {
	var sb strings.Builder
	sb.WriteString(b.name)
	if len(b.state) > 0 {
		sb.WriteRune('[')

		keys := make([]string, 0, len(b.state))
		for key := range b.state {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		first := true
		for _, key := range keys {
			if first {
				first = false
			} else {
				sb.WriteRune(',')
			}
			sb.WriteString(key)
			sb.WriteRune('=')
			sb.WriteString(b.state[key])
		}

		sb.WriteRune(']')
	}
	return sb.String()
}

func (b1 Block) Equals(b2 Block) bool {
	return reflect.DeepEqual(b1, b2)
}

func (b Block) IsAir() bool {
	return b.name == "air"
}

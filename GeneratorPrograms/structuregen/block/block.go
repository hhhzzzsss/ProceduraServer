package block

import (
	"reflect"
	"sort"
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
	dir, ok := b.state["facing"]
	if ok {
		var newState map[string]string
		for key, value := range b.state {
			newState[key] = value
		}
		newState["facing"] = direction.RotateDirectionString(dir, a)
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

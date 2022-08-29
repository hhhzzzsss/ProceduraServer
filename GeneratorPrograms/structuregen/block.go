package structuregen

type Block struct {
	Name  string
	State map[string]string
}

func MakeBlock(name string) Block {
	return Block{
		name,
		make(map[string]string),
	}
}

func (b Block) Rotate(a int) Block {
	if dir, ok := b.State["facing"]; ok {
		b.State["facing"] = RotateDirectionString(dir, a)
	}
	return b
}

func (b *Block) ToString() string {
	str := b.Name
	if len(b.State) > 0 {
		str += "["
		first := true
		for key, value := range b.State {
			if first {
				first = false
			} else {
				str += ","
			}
			str += key + "=" + value
		}
		str += "]"
	}
	return str
}

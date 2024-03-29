package direction

import "github.com/hhhzzzsss/procedura-generator/util"

type Direction int

const (
	West Direction = iota
	North
	East
	South
	Up
	Down
)

var DirectionNames [6]string = [6]string{
	"west",
	"north",
	"east",
	"south",
	"up",
	"down",
}

var DirectionOffsets [6]util.Vec3i = [6]util.Vec3i{
	util.MakeVec3i(-1, 0, 0),
	util.MakeVec3i(0, 0, -1),
	util.MakeVec3i(1, 0, 0),
	util.MakeVec3i(0, 0, 1),
	util.MakeVec3i(0, 1, 0),
	util.MakeVec3i(0, -1, 0),
}

func (d Direction) Rotate(a int) Direction {
	if d < 4 {
		return Direction(((int(d)+a)%4 + 4) % 4)
	} else {
		return d
	}
}

func RotateDirectionString(direction string, a int) string {
	a = (a%4 + 4) % 4
	for i := 0; i < 4; i++ {
		if direction == DirectionNames[i] {
			return DirectionNames[(i+a)%4]
		}
	}
	return direction
}

func RotateVec(vec util.Vec3i, a int) util.Vec3i {
	a = (a%4 + 4) % 4
	if a == 1 {
		return util.MakeVec3i(-vec.Z, vec.Y, vec.X)
	} else if a == 2 {
		return util.MakeVec3i(-vec.X, vec.Y, -vec.Z)
	} else if a == 3 {
		return util.MakeVec3i(vec.Z, vec.Y, -vec.X)
	} else {
		return vec
	}
}

package structuregen

import "github.com/hhhzzzsss/procedura-generator/util"

type Direction int

const (
	North int = iota
	East
	South
	West
	Up
	Down
)

var directionCycle [4]string = [4]string{
	"north",
	"east",
	"south",
	"west",
}

func (d Direction) Rotate(a int) Direction {
	if d < 4 {
		return Direction((int(d) + a) % 4)
	} else {
		return d
	}
}

func RotateDirectionString(direction string, a int) string {
	for i := 0; i < 4; i++ {
		if direction == directionCycle[i] {
			return directionCycle[(i+a)%4]
		}
	}
	return direction
}

func RotateVec(vec util.Vec3i, a int) util.Vec3i {
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

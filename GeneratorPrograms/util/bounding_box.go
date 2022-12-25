package util

// Bounding box with inclusive upper and lower bounds
type BoundingBox struct {
	X1, Y1, Z1 int
	X2, Y2, Z2 int
}

type BoundingRect struct {
	A1, B1 int
	A2, B2 int
}

func (bb BoundingBox) Contains(pos Vec3i) bool {
	return pos.X >= bb.X1 &&
		pos.Y >= bb.Y1 &&
		pos.Z >= bb.Z1 &&
		pos.X <= bb.X2 &&
		pos.Y <= bb.Y2 &&
		pos.Z <= bb.Z2
}

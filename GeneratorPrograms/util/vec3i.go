package util

import "math"

type Vec3i struct {
	X, Y, Z int
}

func MakeVec3i(x, y, z int) Vec3i {
	return Vec3i{x, y, z}
}

func (v1 Vec3i) Add(v2 Vec3i) Vec3i {
	return Vec3i{
		v1.X + v2.X,
		v1.Y + v2.Y,
		v1.Z + v2.Z,
	}
}

func (v1 Vec3i) Sub(v2 Vec3i) Vec3i {
	return Vec3i{
		v1.X - v2.X,
		v1.Y - v2.Y,
		v1.Z - v2.Z,
	}
}

func (v Vec3i) Scale(factor int) Vec3i {
	return Vec3i{
		v.X * factor,
		v.Y * factor,
		v.Z * factor,
	}
}

func (v Vec3i) LengthSquared() int {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3i) Length() float64 {
	return math.Sqrt(float64(v.LengthSquared()))
}

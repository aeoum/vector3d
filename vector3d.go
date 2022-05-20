package vector3d

import (
	"math"
)

const epsilon = 0.000000001

func equalFloat64(a, b float64) bool {
	return math.Abs(a-b) <= epsilon
}

type Vector3D struct {
	x, y, z float64
}

func New(x, y, z float64) *Vector3D {
	return &Vector3D{x, y, z}
}

func (v *Vector3D) Coordinates() (float64, float64, float64) {
	return v.x, v.y, v.z
}

func (v *Vector3D) Norm() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v *Vector3D) Normalize() *Vector3D {
	if v.IsZeroVector() {
		return v
	}

	norm := v.Norm()
	v.x = v.x / norm
	v.y = v.y / norm
	v.z = v.z / norm

	return v
}

func (v *Vector3D) IsNormalized() bool {
	return v.Norm() == 1
}

func (v *Vector3D) IsZeroVector() bool {
	return v.x == 0 && v.y == 0 && v.z == 0
}

func (v *Vector3D) DotProduct(v2 *Vector3D) float64 {
	return v.x*v2.x + v.y*v2.y + v.z*v2.z
}

func (v *Vector3D) CrossProduct(v2 *Vector3D) Vector3D {
	return Vector3D{
		x: v.y*v2.z - v.z*v2.y,
		y: v.z*v2.x - v.x*v2.z,
		z: v.x*v2.y - v.y*v2.x,
	}
}

func (v *Vector3D) IsEqual(v2 *Vector3D) bool {
	return equalFloat64(v.x, v2.x) && equalFloat64(v.y, v2.y) && equalFloat64(v.z, v2.z)
}

func (v *Vector3D) Scale(a float64) *Vector3D {
	v.x = a * v.x
	v.y = a * v.y
	v.z = a * v.z

	return v
}

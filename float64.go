package vector3d

import (
	"math"
)

const epsilon = 0.000000001

func equalFloat64(a, b float64) bool {
	return math.Abs(a-b) <= epsilon
}

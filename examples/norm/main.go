package main

import (
	"fmt"
	"github.com/zeta-f/vector3d"
)

func main() {
	vec := vector3d.New(1, 2, 3)

	fmt.Printf("The Euclidean norm of vec is %v\n", vec.Norm())
}

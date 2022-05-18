package main

import (
	"fmt"
	"github.com/zeta-f/vector3d"
)

func main() {
	vec := vector3d.New(1, 2, 3)
	fmt.Printf("The Euclidean norm (before normalization) %v\n", vec.Norm())
	fmt.Printf("The Euclidean norm (after normalization) %v\n", vec.Normalize().Norm())
}

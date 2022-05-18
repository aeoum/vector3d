package main

import (
	"fmt"
	"github.com/zeta-f/vector3d"
)

func main() {
	v1 := vector3d.New(1, 2, 3)
	v2 := vector3d.New(3, 2, 1)

	fmt.Printf("The dot product of v1 and v2 is %v\n", v1.DotProduct(v2))
}

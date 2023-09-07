package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	//d := 0.01

	for i := 0; z > 0.1; i++ {
		z -= (z*z - x) / (2 * z)

		fmt.Println(i, z, z*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}

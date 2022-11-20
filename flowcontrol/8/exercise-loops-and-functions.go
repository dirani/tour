package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	const p = float64(1.0e-15)
	fmt.Println("z =", z, "p =", p)
	var diff float64 = x
	for diff > p {
		diff = math.Abs(x - z*z)
		z -= (z*z - x) / (2 * z)
		fmt.Println("z =", z, "diff =", diff)
	}
	fmt.Println("diff =", diff)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}

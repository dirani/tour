package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	// initial guess for z = sqrt(x)
	z := float64(1)
	// tolerance p
	const p = float64(1.0e-15)
	// debug
	fmt.Println("z =", z, "\t\t\t", "p =", p)
	// difference between x and current z^2
	var diff float64 = x
	for diff > p {
		diff = math.Abs(x - z*z)
		z -= (z*z - x) / (2 * z)
		fmt.Println("z =", z, "\t", "diff =", diff)
	}
	// debug
	fmt.Print("diff =\t\t", diff, "\n")
	return z
}

func main() {
	fmt.Print("sqrt(", 2, ") = \t", sqrt(2), "\n")
	fmt.Print("math.Sqrt(", 2, ") = \t", math.Sqrt(2), "\n")
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z, f)
	fmt.Printf("%v %v %v %v\n", x, y, z, f)
	fmt.Printf("%T %T %T %T\n", x, y, z, f)
}

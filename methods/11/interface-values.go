package main

import (
	"fmt"
	"math"
)

type I interface { // I has method M()
	M()
}

// ***
type T struct {
	S string
}

func (t *T) M() { // M() is a method with a type T receiver
	fmt.Println(t.S) // type T implements interface I because it has M()
}

// ***

// ***
type F float64

func (f F) M() { // F implements I thorugh M()
	fmt.Println(f)
}

// ***

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

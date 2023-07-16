package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type Vertex3D struct {
	X, Y, Z float64
}

func (u Vertex3D) Scale3D(f float64) Vertex3D {
	u.X = u.X * f
	u.Y = u.Y * f
	u.Z = u.Z * f
	return u
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	u := Vertex3D{3, 4, 5}
	w := u.Scale3D(10)
	fmt.Println(w)
}

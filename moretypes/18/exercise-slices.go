package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var sx [][]uint8
	for x := 0; x < dx; x++ {
		var sy []uint8
		for y := 0; y < dy; y++ {
			//sy = append(sy, uint8((x+y)/2))
			//sy = append(sy, uint8(x*y))
			sy = append(sy, uint8(x^y))

		}
		sx = append(sx, sy)
	}
	return sx
}

func main() {
	pic.Show(Pic)
}

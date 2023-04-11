package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i := 0
	f0 := 0
	f1 := 1
	f := 0
	return func() int {
		switch i {
		case 0:
			{
				i++
				return 0
			}

		case 1:
			{
				i++
				return 1
			}

		default:
			{
				i++
				f = f0 + f1
				f0 = f1
				f1 = f
			}
			return f
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

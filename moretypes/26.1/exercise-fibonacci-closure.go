package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f0 := 0
	f1 := 1
	i := 0
	fib := 0
	return func() int {
		if i == 0 {
			i++
			return 0
		}
		if i == 1 {
			i++
			return 1
		}
		fib = f0 + f1
		f0 = f1
		f1 = fib

		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

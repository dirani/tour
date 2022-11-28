package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		fmt.Print("i = ", i, "\t")
		pow[i] = 1 << uint(i) // == 2**i
		fmt.Printf("pow[%d] = %d\n", i, pow[i])
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

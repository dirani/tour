package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3}
	s = s[0:1]
	fmt.Println(s, len(s), cap(s))
	var t []int
	s = t
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	sfield := strings.Fields(s)
	fmt.Println("line 12, s.Fields =", sfield) // DEBUG
	m := make(map[string]int)
	for i := 0; i < len(sfield); i++ {
		fmt.Printf("sfield[%v] %v\n", i, sfield[i]) // DEBUG
		m[sfield[i]]++
	}

	//return map[string]int{"x": 1}
	fmt.Println("line 20, m: ", m) // DEBUG
	return m
}

func main() {
	wc.Test(WordCount)
}

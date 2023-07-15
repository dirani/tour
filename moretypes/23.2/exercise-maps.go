package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {

	m := make(map[string]int)
	words := strings.Fields(s)
	fmt.Println("result: ", words)
	fmt.Println("len(s): ", len(s))
	fmt.Println("len(words): ", len(words))

	for _, v := range words {
		m[v] += 1
	}

	//return map[string]int{"x": 1}
	return m
}

func main() {
	//wc.Test(WordCount)
	fmt.Println("WordCount: ", WordCount("I am learning Go! I'm really learning it!"))
}

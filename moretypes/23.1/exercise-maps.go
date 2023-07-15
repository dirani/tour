package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	sfield := strings.Fields(s)
	//fmt.Println("s.Fields =", sfield)
	m := make(map[string]int)
	for i := 0; i < len(sfield); i++ {
		//fmt.Println(sfield[i])
		m[sfield[i]]++
	}

	//return map[string]int{"x": 1}
	return m
}

func main() {
	wc.Test(WordCount)
}

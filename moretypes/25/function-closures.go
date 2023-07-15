package main

import "fmt"

func adder() func(int) int {
	sum := 0
	fmt.Println("sum := 0") // DEBUG
	return func(x int) int {
		fmt.Println("x: ", x)     //DEBUG
		fmt.Println("sum: ", sum) // DEBUG
		sum += x
		return sum
	}
}

func main() {
	//pos, neg := adder(), adder()
	neg := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			//pos(i),
			neg(-2 * i),
		)
	}
}

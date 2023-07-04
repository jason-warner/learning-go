package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		println("x:", x)
		println("sum:", sum)
		sum += x
		return sum
	}
}

func main() {
	sum := adder()
	for i := 0; i < 10; i++ {
		fmt.Println("result:", sum(i))
	}
}

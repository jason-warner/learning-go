package main

import (
	"fmt"
	"math"
)

const name = "Jason"

func greeting(name string) string {
	return "Wassup " + name
}

func getSum(num1 int, num2 float32) int {
	return num1 + int(math.Ceil(float64(num2)))
}

func main() {
	fmt.Println(greeting(name))

	fmt.Println("You are ", getSum(10, 16.3), " years old.")
}

package main

import "fmt"

func main() {

	a := 5
	b := &a

	fmt.Printf("A: %d\n", a)
	fmt.Printf("B: %d\n", b)

	fmt.Println("A: ", a)
	fmt.Println("B: ", b)

	fmt.Printf("A: %T\n", a)
	fmt.Printf("B: %T\n", b)

	fmt.Println(*b)

	*b = 2993

	fmt.Println(a)
}

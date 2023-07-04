package main

import "fmt"

func main() {

	name, age := "Jason", 26.3

	fmt.Println(name)
	fmt.Println(age)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
}

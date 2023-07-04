package main

import "fmt"

func main() {

	// var fruitArr [2]string

	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// fruitArr := [2]string{"Apple", "Orange"}

	fruitArr := []string{"Apple", "Orange", "Banana"}
	fruitSlice := fruitArr[1:2]
	fmt.Println(fruitArr)
	fmt.Println(fruitSlice)

}

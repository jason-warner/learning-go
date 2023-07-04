package main

import "fmt"

func main() {

	ids := []int{
		33,
		55,
		68,
		983,
		3,
	}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0

	for _, id := range ids {
		sum += id
	}

	emails := map[string]string{
		"Bob":   "bob@mail.com",
		"Jason": "jason@mail.com",
	}

	println("sum:", sum)

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
}

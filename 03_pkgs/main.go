package main

import (
	"fmt"
	"math"

	"github.com/jason-warner/learning-go/03_pkgs/strutil"
)

func main() {

	name, age := "Jason", 26.3

	fmt.Println(name)
	fmt.Println(math.Floor(age))
	fmt.Println(math.Ceil(age))

	result := strutil.Reverse(name)
	fmt.Println(result)

}

package main

import (
	"fmt"
)

type IlessThan struct {
	msg string
	x   int
	y   int
}

func lessThan(x int, y int) IlessThan {
	if x < y {
		return IlessThan{
			msg: "%d is less that %d\n",
			x:   x,
			y:   y,
		}
	} else if x > y {
		return IlessThan{
			msg: "%d is not less that %d\n",
			x:   x,
			y:   y,
		}
	} else {
		return IlessThan{
			msg: "%d and %d are equal\n",
			x:   x,
			y:   y,
		}
	}
}

func switcher(color string) string {
	switch color {
	case "red":
		return "color is red"
	case "green":
		return "color is green"
	case "blue":
		return "color is blue"
	default:
		return "color is unknown"
	}
}

func main() {
	answer := lessThan(9, 10)
	fmt.Printf(answer.msg, answer.x, answer.y)
	fmt.Println(switcher("blue !"))
}

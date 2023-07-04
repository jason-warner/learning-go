package main

import (
	"fmt"
	"strconv"
)

//	type Person struct {
//		firstName string
//		lastName  string
//		city      string
//		gender    string
//		age       int
//	}
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// VALUE RECEIVER
func (x Person) greet() string {
	return "Wassup, " + x.firstName + " " + x.lastName + " you are " + strconv.Itoa(x.age) + " years old.\n"
}

// POINTER RECEIVER
func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) gotMarried(spouseLastName string) {
	if p.gender == "Male" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {

	person1 := Person{
		firstName: "Jason",
		lastName:  "Warner",
		city:      "Charlotte",
		gender:    "Male",
		age:       26,
	}

	person2 := Person{
		firstName: "Bailey",
		lastName:  "Ryan",
		city:      "Charlotte",
		gender:    "Female",
		age:       26,
	}
	// person1 := Person{
	// 	"Jason",
	// 	"Warner",
	// 	"Charlotte",
	// 	"Male",
	// 	26,
	// }

	person1.age = person1.age + 2
	fmt.Println(person1)
	fmt.Println(person1.firstName)

	person1.hasBirthday()
	person2.gotMarried(person1.lastName)
	fmt.Print(person1.greet())
	fmt.Print(person2.greet())
}

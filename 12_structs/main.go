package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	firstName, lastName, city string // shorter way
	age                       int
}

// A method don't go inside a struct, but it can be called from outside

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old."
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getsMarried (pointer receiver)
func (p *Person) getsMarried(spouseLastName string) {
	if p.lastName == spouseLastName {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	fmt.Println("Hello, World!")

	// Init person using struct
	person1 := Person{
		firstName: "James",
		lastName:  "Bond",
		city:      "London",
		age:       32,
	}

	fmt.Println(person1)

	// Alternative
	person2 := Person{
		"James",
		"Bond",
		"London",
		32,
	}

	fmt.Println(person2)
	fmt.Println(person2.firstName)

	person1.age++
	fmt.Println(person1.age)

	fmt.Println(person1.greet())

	person1.hasBirthday()

	fmt.Println(person1.greet())

	person1.getsMarried("Shaken, not stirred")
	fmt.Println(person1.greet())

}

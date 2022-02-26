package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	// Declare and assign
	names := [2]string{"John", "Paul"}
	fmt.Println(names[0])
	fmt.Println(names[1])

	// Slices
	animals := []string{"Dog", "Cat", "Bird", "Fish"}
	fmt.Println(animals)
	fmt.Println(len(animals))
	fmt.Println(animals[1:3])
}

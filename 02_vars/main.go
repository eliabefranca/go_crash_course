package main

import (
	"fmt"
)

// This will give an error because it is not allowed outside of a function
// name := "Eliabo"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	var age = 27
	var isCool = true

	// Shorthand
	name := "Eliabo" // This shorthand cannot be used outside of a function
	phone, email := "999999", "eliabe.fc@gmail.com"

	fmt.Println(name, email, phone, age, isCool)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", email)
	fmt.Printf("%T\n", phone)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
}

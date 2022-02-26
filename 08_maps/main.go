package main

import "fmt"

func main() {
	// Define map

	emails := make(map[string]string)

	// Assign key value pairs
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

	// declare and assign
	contacts := map[string]string{"John": "99954444", "Paul": "99954444"}
	fmt.Println(contacts)
}

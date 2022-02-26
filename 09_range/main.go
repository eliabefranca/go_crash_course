package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	// Range with map
	contacts := map[string]string{"John": "99954444", "Paul": "99954444"}
	for k, v := range contacts {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range contacts {
		fmt.Println("Name: " + k)
	}
}

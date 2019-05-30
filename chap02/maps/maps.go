package main

import (
	"fmt"
	"sort"
)

func main() {
	// Declare a nil map
	var langs map[int]string

	// Initialize map with make function
	langs = make(map[int]string)

	// Add data as key value pairs
	langs[0] = "Python"
	langs[1] = "Go"
	langs[2] = "JavaScript"
	langs[3] = "Rust"
	fmt.Println("langs data :", langs)

	// Initialize map with map expression
	books := map[string]string{
		"Angels & Demons": "Dan Brown",
		"Godan":           "PremChand",
	}

	// Iterate over the elements of map using range
	for key, value := range langs {
		fmt.Printf("langs[%d]: %s\n", key, value)
	}

	// Lookup an element with key
	if author, ok := books["Godan"]; ok {

		fmt.Println("Godan found in books map")
		fmt.Println("Author of Godan:", author)
	} else {
		fmt.Println("Godan  not found in books map")
	}

	// delete an element
	delete(books, "Godan")

	if author, ok := books["Godan"]; ok {

		fmt.Println("Godan found in books map:", ok)
		fmt.Println("Author of Godan:", author)
	} else {
		fmt.Println("Godan  not found in books map")
		fmt.Println("Error:", ok)
	}

	// Iterate over a Map with an order
	// Slice for specifying the order of the map
	var keys []int
	for key := range langs { // empty identifier won;t work for values
		keys = append(keys, key)
	}
	// Ints sort a slice of ints in increasing order
	sort.Ints(keys)
	for key := range keys {
		fmt.Printf("langs[%d]: %s\n", key, langs[key])
	}

}

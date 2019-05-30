package main

import "fmt"

func main() {
	// Declare array
	var numList [5]int
	numList[0] = 1
	numList[1] = 2
	numList[2] = 3
	numList[3] = 4
	numList[4] = 5
	fmt.Println("Value of numList:", numList)

	// Declare and initialize array with array literal
	oddNumList := [5]int{1, 3, 5, 7, 9}
	fmt.Println("Value of oddNumList:", oddNumList)

	// Array litera with ... (no need to specify length)
	evenNumList := [...]int{0, 2, 4, 6, 8, 10}
	fmt.Println("Value of evenNumList:", evenNumList)
	fmt.Println("Length of evenNumList:", len(evenNumList))

	// Initialize values at specific index with array literal
	langs := [5]string{0: "Python", 4: "Rust"}
	langs[1] = "Go"
	langs[2] = "JavaScript"
	fmt.Println("Value of langs:", langs)

	// Iterate over the elents of array
	for i := 0; i < len(langs); i++ {
		fmt.Printf("langs[%d]: %s", i, langs[i])
	}

	// Iterate over the elements of array by using range
	for index, value := range langs {
		fmt.Println("Index:", index, "Value:", value)
	}

	// use blank identifier
	for _, oddNum := range oddNumList {
		fmt.Println(oddNum)
	}

	// while loop simultion
	i := 0
	for i < len(evenNumList) {
		fmt.Println("Index:", i, "Value:", evenNumList[i])
		i++
	}
}

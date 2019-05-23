package main

import "fmt"

// Fumction to calculate sum
func Sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println("Sum:", Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	// Declare slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Provide slice as argument
	fmt.Println("Sum:", Sum(numbers...))
}

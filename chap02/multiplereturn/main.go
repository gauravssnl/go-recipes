package main

import "fmt"

// Function to swap numbers
func Swap(x, y int) (int, int) {
	return y, x
}

func main() {
	a, b := 2, 3
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("After swapping")
	a, b = Swap(b, a)
	fmt.Println("a", a)
	fmt.Println("b", b)
}

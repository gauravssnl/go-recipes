package main

import "fmt"

// Function that accept function as arguments
func SplitValues(f func(sum int) (int, int)) {
	x, y := f(35)
	fmt.Println("x:", x, "y:", y)

	x, y = f(50)
	fmt.Println("x:", x, "y:", y)
}

func main() {
	a, b := 5, 8
	// Declare an anonymous function
	// Anonymous function can access variables from outer scope
	fn := func(sum int) (int, int) {
		x := sum * a / b
		y := sum - x
		return x, y
	}

	// Passing function as argument
	SplitValues(fn)

	x, y := fn(50)
	fmt.Println("x:", x, "y:", y)
}

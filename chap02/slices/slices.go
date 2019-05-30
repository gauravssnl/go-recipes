package main

import "fmt"

func main() {
	// Declaring nil slice
	var x []int
	fmt.Println("x Length:", len(x), "Capacity:", cap(x))
	// A slice must be initialized before initializing values

	// Initialize slice by using make()
	y := make([]int, 3, 5) // 3 is length, 5 is capacity
	fmt.Println("y Length:", len(y), "Capacity:", cap(y))
	fmt.Println("y value:", y)
	y = make([]int, 3) // capacity defaults to length value
	fmt.Println("y Length:", len(y), "Capacity:", cap(y))

	// Creating slices using slice literal
	z := []int{10, 20, 30} // length = 3, capacity = 3
	fmt.Println("z Length:", len(z), "Capacity:", cap(z))
	fmt.Println("z value:", z)

	z = []int{0: 10, 2: 40} // length = 3, capacity = 3 by using highest index
	fmt.Println("z Length:", len(z), "Capacity:", cap(z))
	fmt.Println("z value:", z)

	// create an empty slice
	// empty slice & nil slice are differernt
	z = []int{}
	fmt.Println("z Length:", len(z), "Capacity:", cap(z))
	fmt.Println("z value:", z)
	fmt.Println(z == nil)
	fmt.Println(x == nil) // x is nil slice

	w := make([]int, 3, 4)
	w[0] = 1
	w[1] = 2
	w[2] = 3
	//w[3] = 4 // this will cause runtime error: index out of range (3 is the capacity)
	fmt.Println("w value:", w)
}

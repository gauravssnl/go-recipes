package main

import "fmt"

func main() {
	// Enarging slices with copy function
	x := []int{10, 20, 30}
	fmt.Println("Length of x:", len(x), "and capacity:", cap(x))
	y := make([]int, 5, 10)
	fmt.Println("Length of y:", len(y), "and capacity:", cap(y))
	copy(y, x)
	fmt.Println("Length of y:", len(y), "and capacity:", cap(y))
	fmt.Println("y value after copying x value: ", y)
	y[3] = 40
	y[4] = 50
	fmt.Println("y value after appending values: ", y)

	// Enlarge a slice with append function
	z := make([]int, 2, 5)
	z[0] = 10
	z[1] = 20
	fmt.Println("Slice z:", z)
	fmt.Println("Slice z length:", len(z), "capacity:", cap(z))
	z = append(z, 30, 40, 50)
	fmt.Println("Slice z after appending:", z)
	fmt.Println("Slice z length:", len(z), "capacity:", cap(z))
	z = append(z, 60)
	fmt.Println("Slice z after appending:", z)
	fmt.Println("Slice z length:", len(z), "capacity:", cap(z))

	// Appending data to nil slice
	var num []int
	num = append(num, 1, 2, 3, 4)
	fmt.Println(num, len(num), cap(num))
}

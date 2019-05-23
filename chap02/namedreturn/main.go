package main

import "fmt"

// Function with named return
func Add(x, y int) (result int) {
	result = x + y
	return // naked return
}

func main() {
	fmt.Println(Add(2, 3))
}

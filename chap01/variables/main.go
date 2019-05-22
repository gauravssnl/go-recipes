package main

import "fmt"

// Decalartion & Assignemnt

// Title : Declare constant
const Title = "GOLANG"

// package level variable
var data = "Package Variable"

func main() {
	fmt.Println("Title:", Title)
	fmt.Println("data:", data)

	// Declare variables
	var s1 string
	s1 = "Test"
	fmt.Println("s1:", s1)
	// Declare and initialize multiple variables
	var fname, lname = "Boy", "Gaurav"
	fmt.Printf("Name: %s %s", fname, lname)
	fmt.Println() // Prints an empty line
	// short value decalaration (only inside functions)
	x, y := 5, 3
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	x, y = y, x
	fmt.Println("After Swapping x & y")
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	// Find default value of different data types
	var s string
	var i int
	var b bool
	var f float32
	// Declare array
	var arr []string
	var c chan []int // channel
	fmt.Println("s:", s)
	fmt.Println("i:", i)
	fmt.Println("b:", b)
	fmt.Println("f:", f)
	fmt.Println("arr", arr)
	fmt.Println("c:", c)
}

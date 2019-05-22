package main

import "fmt"

// This code will compile fine, but by using govet we can detect the bug

func main() {
	fmt.Printf("Float %d", 2.5)
}

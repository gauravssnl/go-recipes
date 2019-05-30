package main

import "fmt"

func panicRecover() {
	defer fmt.Println("Deferred call - 1")
	// expression in defer must be function call
	defer func() {
		fmt.Println("Deferred call - 2")
		if e := recover(); e != nil {
			// e is the value passed to the panic
			fmt.Println("Recover with: ", e)
		}
	}()
	panic("Just panicking for the sake of example")
	fmt.Println("This will never be executed")
}

func main() {
	fmt.Println("Starting to Panic")
	panicRecover()
	fmt.Println("Program regains controll after panic recovery")
}

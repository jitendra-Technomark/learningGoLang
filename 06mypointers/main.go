package main

import (
	"fmt"
)

// Pointer is not just the copy of the value but it use the exact value for any operation we do with that pointer value.
func main() {
	fmt.Println("Pointers in Golang")

	// var ptr *int
	// fmt.Println("Value of Pointer is .", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of Pointer is .", ptr)
	fmt.Println("Value of Pointer is .", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber)
}

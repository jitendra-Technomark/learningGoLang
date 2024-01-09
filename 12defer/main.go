package main

import "fmt"

func main() {
	fmt.Println("Defer Examples")
	defer fmt.Println("Hello World")

	myDefer()
}

// defer works with LIFO functionality
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

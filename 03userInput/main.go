package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating:")

	// comma ok || err err syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for reading, ", input)
	fmt.Printf("Type of this rating is %T", input)

	// secondInput, _ := reader.ReadString('\n')
	// fmt.Println("The second input is, ", secondInput)
	// fmt.Printf("Type of this rating is %T", secondInput)
}

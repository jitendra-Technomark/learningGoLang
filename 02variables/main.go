package main

import "fmt"

func main() {
	var username string = "Jitendra"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("variable is of type: %T \n", smallValue)

	var smallFloat float32 = 255.45545454545454545454545454545454545
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	var bigFloat float64 = 255.45545454545454545454545454545454545
	fmt.Println(bigFloat)
	fmt.Printf("variable is of type: %T \n", bigFloat)

	// default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	var defaultStringVariable string
	fmt.Println(defaultStringVariable)
	fmt.Printf("variable is of type: %T \n", defaultStringVariable)
}

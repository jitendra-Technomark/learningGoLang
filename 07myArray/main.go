package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays in Golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Pineapple"
	// fruitList[0] = "Peach"

	fmt.Println("fruit list is ", fruitList)
	fmt.Println("fruit list is ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "cabbage"}
	fmt.Println("veg list is ", vegList)
	fmt.Println("veg list is ", len(vegList))

}

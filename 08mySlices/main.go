package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in Golang")

	var fruitList = []string{"Apple", "grapes", "Peach"}
	fmt.Printf("Type of the fruitlist is %T\n ", fruitList)

	fruitList = append(fruitList, "Mango", "orange")
	fmt.Println(fruitList)

	// fruitList = append(fruitList[1:])
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 245
	highScores[1] = 125
	highScores[2] = 354
	highScores[3] = 275
	// highScores[5] = 275

	highScores = append(highScores, 555, 666)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	// remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "Golang"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}

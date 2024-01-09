package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a file."

	file, err := os.Create("./myTextFile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNillErr(err)

	length, err := io.WriteString(file, content)

	checkNillErr(err)

	fmt.Println("length is: ", length)
	defer file.Close()

	readFile("./myTextFile.txt")
}

func readFile(fileName string) {
	databyte, err := os.ReadFile(fileName)

	checkNillErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study to golang")

	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2023, time.December, 01, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}

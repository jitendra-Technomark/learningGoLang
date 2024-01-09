package main

import "fmt"

func main() {
	fmt.Println("Struct in golang.")

	Jitendra := User{"Jitendra", "jitendra123@gmail.com", true, 24}
	fmt.Println(Jitendra)
	fmt.Println(Jitendra.Email)

	fmt.Printf("Jitendra's details are %+v\n", Jitendra)

	fmt.Printf("Name is %v and Email is %v", Jitendra.Name, Jitendra.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

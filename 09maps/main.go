package main

import "fmt"

func main() {

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["TS"] = "TypeScript"
	languages["PY"] = "Python"

	fmt.Println(languages)
	fmt.Println("JS shorts for", languages["JS"])

	delete(languages, "TS")
	fmt.Println(languages)

	// loops in golang
	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("value is %v\n", value)
	}

}

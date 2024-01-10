package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	encodeJson()
	decodeJson()
}

func encodeJson() {

	youtubeCourses := []course{
		{"ReactJS Bootcamp", 299, "Youtube", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 299, "Youtube", "bcd123", []string{"web-dev", "fullstack"}},
		{"Golang Bootcamp", 399, "Youtube", "hig123", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(youtubeCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
	// fmt.Println(string(finalJson))
}

func decodeJson() {
	jsonDataFromWeb := []byte(`
	
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "Youtube",
		"tags": ["web-dev", "js"]
	}
	
	`)

	var youtubeCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was Valid!")
		json.Unmarshal(jsonDataFromWeb, &youtubeCourse)

		fmt.Printf("%#v\n", youtubeCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID!")
	}

	// Some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and Type is %T\n", k, v, v)
	}
}

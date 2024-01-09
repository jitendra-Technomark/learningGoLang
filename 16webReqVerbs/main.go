package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	performGetRequest()

	performPostJsonRequest()

	performPostFormEncodedRequest()
}

func performGetRequest() {
	const myURL string = "http://localhost:8000/get"

	response, err := http.Get(myURL)

	if err != nil {
		panic(err)
	}

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	var responseString strings.Builder

	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())

	fmt.Println(string(content))

	defer response.Body.Close()
}

func performPostJsonRequest() {
	const myURL = "http://localhost:8000/post"

	// fake JSON payload
	requestBody := strings.NewReader(`
		{
			"coursename":"learn go lang",
			"price": 0,
			"plateform":"youtube.com"
		}
	`)

	response, err := http.Post(myURL, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

	defer response.Body.Close()
}

func performPostFormEncodedRequest() {
	const myURL string = "http://localhost:8000/postform"

	// formdata
	data := url.Values{}
	data.Add("firstname", "Jitendra")
	data.Add("lastname", "suthar")
	data.Add("email", "jitendra123@gmail.com")

	response, err := http.PostForm(myURL, data)

	if err != nil {
		panic(err)
	}

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

	defer response.Body.Close()

}

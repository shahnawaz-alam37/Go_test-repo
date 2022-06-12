package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("simple get request...")
}

func PerformGetrequest() {
	const myurl string = "http://localhost:3000/get"

	responce, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()
}

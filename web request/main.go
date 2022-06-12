package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("hlo")
	responce, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("responce is of type %T", responce)
	defer responce.Body.Close()

	databytes, err := ioutil.ReadAll(responce.Body)
	if err != nil {
		panic(err)
	}
	demo := string(databytes)
	fmt.Println(demo)
}

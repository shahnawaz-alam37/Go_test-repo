package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://loc.dev:3000/learn?coursename=reactjs&paymentid=hsdf131hjsdf"

func main() {
	fmt.Println("welcome to handle URLs")
	responce, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}

	fmt.Println(responce.Scheme)
	fmt.Println(responce.Host)
	fmt.Println(responce.Path)
	fmt.Println(responce.RawQuery)
	fmt.Println(responce.Port())
	qparams := responce.Query()
	fmt.Printf("the params are %T\n", qparams)
	fmt.Println(qparams["coursename"])
	for _, val := range qparams {
		fmt.Println(val)
	}
	createurl := &url.URL{
		Scheme: "https",
		Host:   "google.com",
		Path:   "/doodles",
	}
	anostr := createurl.String()
	fmt.Println(anostr)
}

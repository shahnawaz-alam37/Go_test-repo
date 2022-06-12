package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("simple get request...")
	//PerformGetrequest()
	//Performpostjsonrequest()
	PerformPostFormRequest()
}

func PerformGetrequest() {  
	const myurl string = "http://localhost:3000/get"

	responce, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()
	fmt.Println("status code:",responce.StatusCode)
	fmt.Println("content length:",responce.ContentLength)
	content , _ := ioutil.ReadAll(responce.Body)
	var responsestring strings.Builder					/*this strings package is optional*/
	bytecount , _ := responsestring.Write(content)
	fmt.Println(bytecount)
	fmt.Println(responsestring.String())
	//fmt.Println(string(content))
}

func Performpostjsonrequest()  {
	const myurl = "http://localhost:3000/post"

	requestBody := strings.NewReader(`
		{
			"name":"shahnawaz",
			"pin":37,
			"collage":"QQGPT"
		}
	`)
	response, err := http.Post(myurl,"application/json",requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func PerformPostFormRequest()  {
	const myurl string = "http://localhost:3000/postform"

	data := url.Values{}
	data.Add("firstname","shahnawaz")
	data.Add("lastname","alam")
	data.Add("collage","QQGPT")

	response , err := http.PostForm(myurl,data,)

	if err != nil {
		panic(err)
	}
	content, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Println(string(content))
}



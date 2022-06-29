package main

import (
	"fmt"
	"net/http"
	"sync"
)
var wg sync.WaitGroup

func main()  {
	fmt.Println("go rotines")
	var urls = []string{
		"https://github.com",
		"https://google.com",
		"https://lco.dev",
		"https://youtube.com",
	}
	for _, v := range urls {
		go getstatuscode(v)
		wg.Add(1)
	}
	wg.Wait()
}

func getstatuscode(state string)  {
	defer wg.Done()
	res , err := http.Get(state)
	if err != nil {
		fmt.Println("error in status code")
	}else{
		fmt.Printf("%d is the status code of %s\n",res.StatusCode,state)
	}

}
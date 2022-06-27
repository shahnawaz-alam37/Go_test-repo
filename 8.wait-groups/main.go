package main

import (
	"fmt"
	"net/http"
	
	"sync"
)

var signal = []string{"test"}
var wg sync.WaitGroup
var mux sync.Mutex


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
	fmt.Println(signal)
}

func getstatuscode(state string)  {
	defer wg.Done()
	res , err := http.Get(state)
	if err != nil {
		fmt.Println("error in status code")
	}else{
		mux.Lock()
		signal = append(signal, state)
		mux.Unlock()
		
		fmt.Printf("%d is the status code of %s\n",res.StatusCode,state)
	}

}
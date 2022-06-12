package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("hello world")
	x := make(map[int]string)
	x[1] = "shah"
	x[2] = "alam"
	x[3] = " "
	fmt.Println(x)
	fmt.Printf("enter your name:")
	fmt.Scanln(&name)

}

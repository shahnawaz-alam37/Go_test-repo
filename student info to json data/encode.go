package main

import (
	"encoding/json"
	"fmt"
)
type info struct{
	Name string    `json:"student name"`
	Pin int		   `json:"pin"`
	Collage string `json:"collage-name"`
	Branch string  `json:"branch"`
}
func main()  {
	fmt.Println("")
	Encodejson()
}

func Encodejson()  {

	var v1 info
	fmt.Printf("enter your name:")
	fmt.Scanln(&v1.Name)

	fmt.Printf("enter your pin:")
	fmt.Scanln(&v1.Pin)

	fmt.Printf("enter your collage name:")
	fmt.Scanln(&v1.Collage)

	fmt.Printf("enter your branch name:")
	fmt.Scanln(&v1.Branch)
	in := []info{
		{v1.Name,v1.Pin,v1.Branch,v1.Collage},
	}
	encode, err := json.MarshalIndent(in,"","\t")
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n",encode)
}	
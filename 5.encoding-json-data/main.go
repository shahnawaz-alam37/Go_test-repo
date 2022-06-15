package main

import (
	"encoding/json"
	"fmt"
)

type course struct{
	Name 	 string	`json:"name"`
	Price 	 int	
	Platform string	`json:"website"`
	Password string	`json:"-"`
	Tags 	 []string `json:"tags,omitempty"`
}

func main()  {
	fmt.Println("Json data...")
	Encodejson()
}

func Encodejson()  {
	lcoCourses := []course{
		{"reactjs",299,"online","1234",[]string{"webdev","js"}},
		{"vuejs",199,"online","4321",[]string{"webdev","js"}},
		{"angularjs",99,"online","1111",nil},	
	}
	//packing data as json data
	finaljson, err := json.MarshalIndent(lcoCourses,"","\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n",finaljson)

}
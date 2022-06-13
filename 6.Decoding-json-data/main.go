package main

import (
	
	"encoding/json"
	"fmt"
)

//structure to store data
type course struct{
	Name 	 string	`json:"name"`
	Price 	 int	
	Platform string	`json:"website"`
	Password string	`json:"-"`
	Tags 	 []string `json:"tags,omitempty"`
}


func main()  {
	fmt.Println("decoding json data...")
	decodingjson()
}

func decodingjson()  {
	jsondatafromweb := []byte(`
	{
		"name": "reactjs",
		"Price": 299,
		"website": "online",
		"tags": ["webdev","js"]
	}
	`)
	var coreinfo course
	checkvalid := json.Valid(jsondatafromweb)
	if checkvalid {
		fmt.Println("JSON WAS VALID")
		json.Unmarshal(jsondatafromweb,&coreinfo)
		fmt.Printf("%#v\n",coreinfo)
	}else{
		fmt.Println("json was not valid")
	}
}
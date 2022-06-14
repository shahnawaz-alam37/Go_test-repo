package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	
)

//fake database
type Info struct{
	Name string  `json:"name"`
	Pin string   `json:"pin"`
}

//fake database

var ob []Info

func main()  {
	fmt.Println("launching the api")
	
	//seeding
	ob = append(ob, Info{Name: "shahnawaz",Pin: "37"})
	ob = append(ob, Info{Name: "Rasool",Pin: "7"})

	//routing 
	r :=mux.NewRouter()
	r.HandleFunc("/",ServerHome).Methods("GET")
	r.HandleFunc("/infor",Getallinfo).Methods("GET")
	r.HandleFunc("/info/{pin}",Getoneinfo).Methods("GET")
	r.HandleFunc("/info",postsomedata).Methods("POST")

	//listening to a port
	log.Fatal(http.ListenAndServe(":4000",r))

	fmt.Println("creating a simple api")
}

//controllers
func ServerHome(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("<h1>another simple api</h1>"))
}

//printing all the information
func Getallinfo(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("printing all possible info")
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(ob)//passing the fake db
}

//getting one specific info
func Getoneinfo(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("printing the requested info")
	w.Header().Set("Content-type","application/json")

	//get the pin first
	params := mux.Vars(r)

	//looping through the database
	for _, obj := range ob {
		if obj.Pin == params["pin"] {
			json.NewEncoder(w).Encode(obj)
			return
		}
	}
	json.NewEncoder(w).Encode("no data found")	
}

//posting som data
func postsomedata(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("inserting new data")
	w.Header().Set("Content-type","application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("body is empty")
	}

	var obj Info
	_ = json.NewDecoder(r.Body).Decode(&obj)
	ob = append(ob,obj)
	json.NewEncoder(w).Encode(ob)
}


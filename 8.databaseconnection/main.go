package main

import (
	"database/sql"
	"fmt"
	"log"
	"encoding/json"
	_ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "abc123"
  dbname   = "postgres"
)

type info struct{
	Name string `json:name`
	Pin int `json:pin`
	Address string  `json:address`
}
type variable[]info
var slice info



func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()






  quri := "select * from emp2"
  data , err := db.Query(quri)
  if err != nil {
	log.Fatal(err)
  }
var vari variable 
  for data.Next(){
 	 var perinfo info
 	 err = data.Scan(&slice.Name,&slice.Pin,&slice.Address)
  	if err != nil {
		log.Fatal(err)
  	}
  	vari = append(vari, perinfo)
}
	

	in := []info{
		{slice.Name,slice.Pin,slice.Address},
	}
	dat, err := json.MarshalIndent(in,"","\t")
	if err != nil {
		panic(err)
	  }
	  fmt.Printf("%s\n",dat)


  err = db.Ping()
  if err != nil {
    panic(err)
  }else{
	fmt.Println("connection successfull")
  }
}
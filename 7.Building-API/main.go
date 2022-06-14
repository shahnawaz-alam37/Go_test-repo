package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"math/rand"
	"time"
	"strconv"
	"github.com/gorilla/mux"
)
type Course struct{
	CourseId string `json:"course Id"`
	CourseName string `json:"course name"`
	CoursePrice int `json:"price"`
	Author 		*Author `json:"author"`
}
type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}
//fake database 
var courses []Course

//middleware or helpre - file
func (c *Course) Isempty() bool{
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}


func main(){
	fmt.Println("simple api")
}

//controllers
//serve home route
func serverHome(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("<h1>welcome to my first api</h1>"))
}

func getallcourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOnecourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("get one course")
	w.Header().Set("Content-Type","application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through courses,find the matching id and return the responce
	for _,course := range courses{
		if course.CourseId==params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("create one course")
	w.Header().Set("Content-Type","application/json")

	//what if body is emp
	if r.Body == nil {
		json.NewEncoder(w).Encode("body is empty")
	}
	//what if json is empty {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.Isempty() {
		json.NewEncoder(w).Encode("no data inside the json")
		return
	}
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Student data structure
type Student struct {
	Name  string `json:"Name"`
	Grade string `json:"Grade"`
}

// Students Array
var students []Student

func main() {
	students = []Student{
		Student{Name: "David", Grade: "A+"},
		Student{Name: "Khosro", Grade: "F---"},
	}
	handleRequests()
}

//func to handle all network requests
func handleRequests() {

	muxRouter := mux.NewRouter().StrictSlash(true)

	muxRouter.HandleFunc("/", homePage)
	muxRouter.HandleFunc("/students", returnAllStudents).Methods("GET")
	muxRouter.HandleFunc("/students", createNewStudent).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", muxRouter))
}

//func to display the homepage
// * makes it a "pointer" to the type
func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func createNewStudent(writer http.ResponseWriter, request *http.Request) {
	//ioutil returns 2 variables,
	// first is the body,
	//second is the error message (if any)
	//GoLang has to use ALL variables declared so we used
	//An underscore to tell the compiler we don't need the error
	reqBody, _ := ioutil.ReadAll(request.Body)
	var student Student
	json.Unmarshal(reqBody, &student)
	// update our global students array to include
	// our new student
	students = append(students, student)

	json.NewEncoder(writer).Encode(student)

}

//function to run when /students endpoint is hit
func returnAllStudents(writer http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllStudents")
	json.NewEncoder(writer).Encode(&students)
}

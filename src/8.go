
// MySchoolProject is a simple web application that displays a list of students and their grades.
package main

import (
	"fmt"
	"log"
	"net/http"
)

type Student struct {
	Name  string
	Grade int
}

var Students = []Student{
	{"Alice", 90},
	{"Bob", 75},
	{"Charlie", 82},
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to MySchoolProject!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
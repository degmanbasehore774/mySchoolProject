package main

import "fmt"

func main() {
    name := "My School Project"
    course := "A new course for my school projects."
    students := []string{"Alice", "Bob"}
    
    fmt.Println("Hello, ", name)
    fmt.Println("Here are the courses:")
    for _, student := range students {
        fmt.Printf("%s: %s\n", student, course)
    }
}

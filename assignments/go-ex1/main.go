package main

import "fmt"

// Student represents a struct with Name and Course fields.
type Student struct {
	Name   string
	Course string
}

// Describe is a method with a receiver that prints the struct's contents.
func (s Student) Describe() {
	fmt.Printf("Hello! My name is %s and I am studying %s.\n", s.Name, s.Course)
}

func main() {
	me := Student{Name: "Mustapha Oluwatoyin Gali", Course: "Artificial Intelligent Solution at FH Upper Austria"}
	me.Describe()
}

package main

import "fmt"

// 1️⃣ Define the struct
type Student struct {
	Name    string
	Program string
	School  string
}

// 2️⃣ Define a method with receiver
func (s Student) Introduce() {
	fmt.Printf("Hello! My name is %s and I am studying %s at %s.\n",
		s.Name, s.Program, s.School)
}

// 3️⃣ Call the method in main
func main() {
	me := Student{
		Name:    "Mustapha Oluwatoyin Gali",
		Program: "Artificial Intelligent Solution",
		School:  "FH Upper Austria",
	}
	me.Introduce()
}

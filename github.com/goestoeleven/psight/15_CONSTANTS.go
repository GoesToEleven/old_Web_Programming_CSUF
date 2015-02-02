package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func main() {

	var s = Salutation{}
	s.name = "Bob"
	s.greeting = "Hello!"

	fmt.Println(s.name)
	fmt.Println(s.greeting)

}

// we can use a STRUCT like a class

// ***********************************************
// we can also add functions (methods) to user defined types
// ***********************************************

// go is not an OOP language
// the type sustem that exists in go
// -- makes it so you don't need classes
// -- gives you more flexibility b/c you're not constrained by class requirements
// instead of using classes, we have user defined types

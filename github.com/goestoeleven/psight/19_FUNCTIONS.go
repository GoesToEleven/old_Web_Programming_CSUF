package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation) {
	fmt.Println(salutation.name)
	fmt.Println(salutation.greeting)
}

func main() {
	var s = Salutation{}
	s.name = "Bob"
	s.greeting = "Hello!"
	Greet(s)
}

// FUNCTIONS IN GO
// -- multiple return values
// -- use like any other type
// ---- similar to JavaScript
// ------ pass into other functions
// ------ declare them as variables
// ------ return them from functions
// -- function literals
// ---- you can declare functions inside other functions

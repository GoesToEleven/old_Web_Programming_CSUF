package main

import "fmt"

type Contact struct {
	name     string
	greeting string
}

func Greet(contact Contact) {
	fmt.Println(contact.name)
	fmt.Println(contact.greeting)
}

func main() {
	var c = Contact{}
	c.name = "Marcus"
	c.greeting = "Hello!"
	Greet(c)
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

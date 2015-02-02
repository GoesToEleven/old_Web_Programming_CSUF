package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func main() {

	var s = Salutation{"Joe", "Hello!"}

	fmt.Println(s.name)
	fmt.Println(s.greeting)

}

// we can use a STRUCT like a class

// go is not an OOP language
// the type sustem that exists in go
// -- makes it so you don't need classes
// -- gives you more flexibility b/c you're not constrained by class requirements
// instead of using classes, we have user defined types

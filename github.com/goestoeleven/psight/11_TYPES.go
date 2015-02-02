package main

import "fmt"

type Salutation string

func main() {

	var message Salutation = "Hello World!"

	fmt.Println(message)

}

// go is not an OOP language
// the type sustem that exists in go
// -- makes it so you don't need classes
// -- gives you more flexibility b/c you're not constrained by class requirements
// instead of using classes, we have user defined types

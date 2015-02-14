package main

import "fmt"

// TYPE
type contact struct {
	name     string
	greeting string
}

// TYPES CAN CONTAIN METHODS & RETURN
func (c contact) sayHello() string {
	return "METHOD: " + c.name + " says hello."
}

func (c contact) rename(newName string) {
	c.name = newName
}

func main() {
	// TYPES CAN CONTAIN DATA
	var friend = contact{"Marcus", "Hello!"}
	fmt.Println("DATA: " + friend.name)
	fmt.Println("DATA: " + friend.greeting)

	// TYPES CAN USE METHODS
	fmt.Println(friend.sayHello())
	friend.rename("Jenny")
	fmt.Println("DATA: " + friend.name)
	fmt.Println(friend.sayHello())
}

// QUESTION:
// Why isn't the name changing?
/*
answer down below














PASS COPY
func (c contact) rename(newName string) {
  c.name = newName
}

PASS REFERENCE
func (c *contact) rename(newName string) {
  c.name = newName
}





*/

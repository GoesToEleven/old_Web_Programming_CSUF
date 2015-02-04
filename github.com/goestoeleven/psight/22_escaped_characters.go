package main

import "fmt"

type Salutation struct {
	greeting string
	name     string
}

func Greet(salutation Salutation) {
	fmt.Println(CreateMessage(salutation.greeting, salutation.name))
}

func CreateMessage(greeting, name string) (string, string) {
	return greeting + " " + name, "\nHey, " + name + "\n\n"
}

func main() {

	var t = Salutation{"Good to see you,", "Medhi"}
	Greet(t)

	u := Salutation{"Glad you're in class,", "Sushant"}
	Greet(u)

	v := Salutation{}
	v.greeting = "We're learning great things,"
	v.name = "Max"
	Greet(v)
}

// put parentheses around return types and comma separate them
// use comma to separate the multiple returns

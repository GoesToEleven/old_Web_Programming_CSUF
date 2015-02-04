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

	var t = Salutation{"Medhi", "Good to see you!"}
	Greet(t)

	u := Salutation{"Sushant", "Glad you're in class!"}
	Greet(u)

	v := Salutation{}
	v.name = "Max"
	v.greeting = "We're learning great things!"
	Greet(v)

}

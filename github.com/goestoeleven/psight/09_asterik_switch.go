package main

import "fmt"

func main() {

	message := "Hello Go World!"

	var greeting *string = &message

	fmt.Println(message, *greeting)

	// the * says, "Show me what you're pointing towards."
}

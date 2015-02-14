package main

import "fmt"

// FOR
// -- clause
// ---- init; cond; post
// -- condition
// -- range

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// will this, below, run?
	// fmt.Println(i)

}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Let's add some delay to the function using time.Sleep and rand.Intn:
*/

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}

/*
f prints out the numbers from 0 to 10, waiting between 0 and 250 ms
after each one. The goroutines should now run simultaneously.
*/

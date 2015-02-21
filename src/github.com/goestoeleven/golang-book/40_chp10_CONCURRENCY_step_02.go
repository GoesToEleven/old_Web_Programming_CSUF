package main

import "fmt"

/*
Goroutines are lightweight and we can easily create thousands of them.
We can modify our program to run 10 goroutines by doing this:
*/

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
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
You may have noticed that when you run this program it seems to run
the goroutines in order rather than simultaneously. Let's change this
in the next program.
*/

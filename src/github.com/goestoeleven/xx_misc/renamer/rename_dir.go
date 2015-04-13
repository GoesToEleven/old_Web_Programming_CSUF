package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	files, _ := ioutil.ReadDir("/Users/tm002/Desktop/changeme")
	for i, f := range files {
		fmt.Println(f.Name())
		oldfile := ("/Users/tm002/Desktop/changeme/" + f.Name())
		newfile := ("/Users/tm002/Desktop/changeme/" + strconv.Itoa(i) + ".png")
		//        fmt.Println(oldfile)
		//        fmt.Println(newfile)
		err := os.Rename(oldfile, newfile)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

/*
    age := 12

    // will NOT display properly
    fmt.Println(":", string(age))

    // convert int to string
    // will display properly
    fmt.Println(":", strconv.Itoa(age))
}
*/

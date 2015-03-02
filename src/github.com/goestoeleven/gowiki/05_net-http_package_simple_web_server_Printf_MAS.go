package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Fprintf(w, "scheme %s!", r.URL.Scheme)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

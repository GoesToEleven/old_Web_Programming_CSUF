package server

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles(
	"site/includes/document_top.html",
	"site/includes/document_bottom.html",
	"site/anybody/index.html",
	"site/anybody/index2.html",
	"site/assets/stylesheets/main.css"))

func init() {
	http.HandleFunc("/index2.html", augh)
	http.HandleFunc("/", root)
	// http.HandleFunc("/sign", sign)
}

func augh(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index2.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func root(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// func sign(w http.ResponseWriter, r *http.Request) {
// 	err := tmpl.ExecuteTemplate(w, "guestbook.html", r.FormValue("content"))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }



/*
NOTES:
http://golang.org/pkg/html/template/#Must
https://golang.org/doc/effective_go.html#data
*/
package main

import (
	"html/template"
	"net/http"
    "log"
)

var mytemp *template.Template

type userstuff struct {
	Name, Email, Message string
}

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/formhandler", formhandler)
	http.HandleFunc("/link1", linkone)
	http.HandleFunc("/link2", linktwo)
	http.HandleFunc("/link3", linkthree)
	mytemp = template.Must(template.ParseFiles("form.html", "formhandler.html", "link1.html", "link2.html", "link3.html"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := mytemp.ExecuteTemplate(w, "form.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func formhandler(w http.ResponseWriter, r *http.Request) {
	user := userstuff{r.FormValue("name"), r.FormValue("email"), r.FormValue("message")}
	err := mytemp.ExecuteTemplate(w, "formhandler.html", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func linkone(w http.ResponseWriter, r *http.Request) {
    log.Println(r.URL.Query())
    log.Println(r.URL.Query().Get("name"))
    log.Println(r.URL.Query().Get("email"))
    log.Println(r.URL.Query().Get("message"))
    user := userstuff{ r.URL.Query().Get("name"), r.URL.Query().Get("email"), r.URL.Query().Get("message")}
    log.Println(user)
    log.Println(user.Name)
    log.Println(user.Email)
    log.Println(user.Message)
    mappedUser := r.URL.Query()
    log.Println("mapped ",mappedUser)
    log.Println(mappedUser["name"])
    log.Println(mappedUser["email"])
    log.Println(mappedUser["message"])
    log.Println(mappedUser["message"][0])
//    log.Println(mappedUser["message"][1])
//    log.Println(mappedUser["message"][2])
    log.Println(mappedUser["name"][0])
    log.Println(mappedUser["email"][0])
    err := mytemp.ExecuteTemplate(w, "link1.html", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func linktwo(w http.ResponseWriter, r *http.Request) {
    mappedUser := r.URL.Query()
    err := mytemp.ExecuteTemplate(w, "link2.html", mappedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func linkthree(w http.ResponseWriter, r *http.Request) {
    user := userstuff{ r.URL.Query().Get("name"), r.URL.Query().Get("email"), r.URL.Query().Get("message")}
	err := mytemp.ExecuteTemplate(w, "link3.html", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

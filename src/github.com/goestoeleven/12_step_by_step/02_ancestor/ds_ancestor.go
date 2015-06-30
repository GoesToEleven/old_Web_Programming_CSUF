package main

import (
    "fmt"
    "net/http"
    "time"

    "appengine"
    "appengine/datastore"
)

type Employee struct {
    Name     string
    Role     string
    HireDate time.Time
    email    string
}

func init() {
    http.HandleFunc("/", myHandler)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)

    var e1 Employee

    userKey := datastore.NewKey(c, "employee", "david@starship.com", 0, nil)

    if err = datastore.Get(c, userKey, &e1); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Retrieved %q", e1)
}

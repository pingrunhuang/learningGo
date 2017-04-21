package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)


func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func DataIndex(w http.ResponseWriter, r *http.Request) {
  // create fake data
    db := NewDB()
    p := db.getAllPersons()
    if err := json.NewEncoder(w).Encode(*p); err != nil {
        panic(err)
    }
}

func display(w http.ResponseWriter, r *http.Request) {
    db := NewDB()
    vars := mux.Vars(r)
    name := vars["name"]
    person := db.getPerson(name)
    fmt.Fprintln(w, "Data display:", person)
}

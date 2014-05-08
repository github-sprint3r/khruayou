package searchnotfound

import (
    "fmt"
    "net/http"
)

func HomeHandler() {
    http.HandleFunc("/", addHanler)
    http.ListenAndServe(":8080", nil)
}

func addHanler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the home page!")
}
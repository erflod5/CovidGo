package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexRoute)

	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8000", nil)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my webiste")
}

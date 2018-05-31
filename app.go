package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey Welcome to my website your path is: %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", indexHandler)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8000", nil)
}

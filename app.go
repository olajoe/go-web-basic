package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey Welcome to Go! %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}

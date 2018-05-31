package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country := vars["name"]
	club := vars["clubName"]
	fmt.Fprintf(w, "You favorite club is %s in %s", club, country)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/countries/{name}/clubs/{clubName}", indexHandler).Methods("GET")

	http.ListenAndServe(":8000", r)
}

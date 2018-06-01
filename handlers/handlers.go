package handlers

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", homePage).Methods("GET")
	r.HandleFunc("/form", formContact).Methods("POST")
	return r
}

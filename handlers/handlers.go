package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method + ": " + r.URL.Path)
		f(w, r)
	}

}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", logging(homePage)).Methods("GET")
	r.HandleFunc("/form", logging(formContact)).Methods("POST")
	return r
}

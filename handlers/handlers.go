package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type middleware func(http.HandlerFunc) http.HandlerFunc

func logging() middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Method + ": " + r.URL.Path)
			f(w, r)
		}
	}
}

func method(m string) middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			f(w, r)
		}
	}
}

func chain(f http.HandlerFunc, middlewares ...middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", chain(homePage, method("GET"), logging()))
	r.HandleFunc("/form", chain(formContact, method("POST"), logging()))

	return r
}

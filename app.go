package main

import (
	"joe/joe-go-api/handlers"
	"net/http"
)

func main() {

	router := handlers.Router()

	http.ListenAndServe(":8000", router)
}

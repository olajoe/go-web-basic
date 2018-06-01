package main

import (
	"joe/go-web-basic/handlers"
	"net/http"
)

func main() {

	router := handlers.Router()

	http.ListenAndServe(":8000", router)
}

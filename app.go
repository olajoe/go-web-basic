package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type clubDetail struct {
	ClubName       string
	Address        string
	LeagueChampion bool
}

type clubData struct {
	CountryName string
	ClubDetails []clubDetail
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("display.html"))
	data := clubData{
		CountryName: "Italy",
		ClubDetails: []clubDetail{
			{ClubName: "AC Milan", Address: "Milan, Italy", LeagueChampion: true},
			{ClubName: "Atalanta", Address: "Bergamo Italy", LeagueChampion: false},
			{ClubName: "Juventus", Address: "Turin, Italy", LeagueChampion: true},
		},
	}
	tmp.Execute(w, data)

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")

	http.ListenAndServe(":8000", r)
}

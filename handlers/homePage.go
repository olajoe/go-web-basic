package handlers

import (
	"html/template"
	"net/http"
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

func homePage(w http.ResponseWriter, r *http.Request) {
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

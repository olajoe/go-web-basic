package handlers

import (
	"html/template"
	"net/http"
)

type contactDetails struct {
	Email   string
	Subject string
	Message string
}

func formContact(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("templates/forms.html")
	details := contactDetails{
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}

	_ = details
	tmp.Execute(w, struct{ Success bool }{true})
}

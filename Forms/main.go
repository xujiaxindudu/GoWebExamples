package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func handle(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// do something with details
		_ = details

		tmpl.Execute(w, struct{ Success bool }{true})
	}
}
func main() {
	tmpl, err := template.ParseFiles("forms.html")
	if err != nil {
		fmt.Println("parse forms failed,err:", err)
		return
	}
	http.Handle("/", handle(tmpl))
	http.ListenAndServe(":80", nil)
}

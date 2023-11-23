package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>Email: %s</p>", r.FormValue("email"))
	fmt.Fprintf(w, "<p>Password: %s</p>", r.FormValue("password"))
}

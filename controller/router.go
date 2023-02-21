package controller

import (
	"html/template"
	"net/http"
)

// starts main page
func MainPage(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("static/templates/index.html"))
	templ.Execute(w, nil)
}

// handles route for login
func Login(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("static/templates/login.html"))
	templ.Execute(w, nil)
}

// handles route for register
func Register(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("static/templates/register.html"))
	templ.Execute(w, nil)
}

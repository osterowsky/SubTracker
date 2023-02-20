package controller

import (
	"html/template"
	"net/http"
)

// starts main page
func MainPage(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("static/index.html"))
	templ.Execute(w, nil)
}

// handles route for login
func Login(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("static/login.html"))
	templ.Execute(w, nil)
}

// handles route for register
func Register(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("static/register.html"))
	templ.Execute(w, nil)
}

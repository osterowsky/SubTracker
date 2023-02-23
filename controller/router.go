package controller

import (
	"html/template"
	"log"
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
	if r.Method == "GET" {
		templ := template.Must(template.ParseFiles("static/templates/register.html"))
		templ.Execute(w, nil)
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		// Get the form values
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		for err = createUser(username, email, password); err != nil; {
			log.Fatal(err)
		}

		// Redirect the user to the index page
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

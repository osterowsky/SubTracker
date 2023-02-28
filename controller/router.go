package controller

import (
	"crypto/md5"
	"encoding/hex"
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
	if r.Method == "GET" {
		templ := template.Must(template.ParseFiles("static/templates/login.html"))
		templ.Execute(w, nil)
	} else if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		passwordHash := md5.Sum([]byte(password))

		for err := loginUser(w, r, username, hex.EncodeToString(passwordHash[:])); err != nil; {
			log.Fatal(err)
			return
		}

		// Redirect the user to the index page
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

// handles route for register
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templ := template.Must(template.ParseFiles("static/templates/register.html"))
		templ.Execute(w, nil)
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
			return
		}
		// Get the form values
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		for err = createUser(username, email, password); err != nil; {
			log.Fatal(err)
			return
		}

		// Redirect the user to the index page
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Get registers and returns a session for the given name and session store.
	session, _ := Store.Get(r, "session.id")
	// Set the authenticated value on the session to false
	session.Values["authenticated"] = false
	session.Save(r, w)
}

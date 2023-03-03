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

	// if user trigger post, we add sub
	if r.Method == "POST" {
		addSub(w, r)
		return
	}

	data := PageData{
		IsLoggedIn: isLogged(r),
	}
	templ := template.Must(template.ParseFiles("static/templates/layout.html", "static/templates/index.html"))
	err := templ.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templ := template.Must(template.ParseFiles("static/templates/layout.html", "static/templates/login.html"))
		if err := templ.ExecuteTemplate(w, "login.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case "POST":

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		passwordHash := md5.Sum([]byte(password))

		// we try to log user
		if err := loginUser(w, r, username, hex.EncodeToString(passwordHash[:])); err != nil {
			log.Fatal(err)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

func Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templ := template.Must(template.ParseFiles("static/templates/layout.html", "static/templates/register.html"))
		if err := templ.ExecuteTemplate(w, "register.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case "POST":
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
			return
		}

		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		if err := createUser(w, r, username, email, password); err != nil {
			log.Fatal(err)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {

	// Get registers and returns a session for the given name and session store.
	session, _ := Store.Get(r, "session.id")
	session.Values["authenticated"] = false
	session.Save(r, w)

	log.Println("User has been logged out")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

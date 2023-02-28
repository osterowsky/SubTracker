package main

import (
	"net/http"
	c "subtracker/controller"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func main() {

	r := mux.NewRouter()

	// serving css inside static resources
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.HandleFunc("/", c.MainPage)
	r.HandleFunc("/login", c.Login)
	r.HandleFunc("/register", c.Register)
	r.HandleFunc("/logout", c.Logout)

	http.ListenAndServe(":8080", r)
}

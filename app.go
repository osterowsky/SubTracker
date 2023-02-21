package main

import (
	"net/http"

	c "subtracker/controller"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", c.MainPage)
	http.HandleFunc("/login", c.Login)
	http.HandleFunc("/register", c.Register)

	http.ListenAndServe(":8080", nil)
}

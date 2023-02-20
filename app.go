package main

import (
	"net/http"

	c "subtracker/controller"
)

func main() {

	http.HandleFunc("/", c.MainPage)
	http.HandleFunc("/login", c.Login)
	http.HandleFunc("/register", c.Register)

	http.ListenAndServe(":8080", nil)
}

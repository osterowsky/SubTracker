package main

import (
	"database/sql"
	"log"
	"net/http"
	c "subtracker/cmd"

	_ "github.com/lib/pq"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", c.MainPage)
	http.HandleFunc("/login", c.Login)
	http.HandleFunc("/register", c.Register)

	http.ListenAndServe(":8080", nil)
}

var db *sql.DB

func init() {
	var err error

	connStr := "postgres://osternpatryk@localhost/subtracker?sslmode=disable"
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	// checks if we are connected to db
	if err = db.Ping(); err != nil {
		panic(err)
	}

	log.Println("The database is connected")
}

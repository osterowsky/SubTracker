package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error

	connStr := "postgres://osternpatryk@localhost:5432/subtracker?sslmode=disable"
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	// checks if we are connected to db
	if err = DB.Ping(); err != nil {
		panic(err)
	}

	log.Println("The database is connected")
}

package controller

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	d "subtracker/schema/database"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("secret"))

func createUser(username, email, password string) (err error) {

	// Generate the MD5 hash of the password
	passwordHash := md5.Sum([]byte(password))
	defer d.DB.Close()

	// Insert the user into the database
	_, err = d.DB.Exec("INSERT INTO users (username, email, passwordHash) VALUES ($1, $2, $3)", username, email, hex.EncodeToString(passwordHash[:]))
	if err != nil {
		return err
	}
	return nil
}

func loginUser(w http.ResponseWriter, r *http.Request, username, passHash string) (err error) {
	var storedHash string
	row, err := d.DB.Query("SELECT passwordHash FROM users WHERE username = $1;", username)

	if err != nil {
		return err
	}
	defer row.Close()

	for row.Next() {
		for err := row.Scan(&storedHash); err != nil; {
			return err
		}
	}

	if len(storedHash) != 0 {
		// It returns a new session if the sessions doesn't exist
		session, _ := Store.Get(r, "session.id")
		if storedHash == passHash {
			session.Values["authenticated"] = true
			// Saves all sessions used during the current request
			session.Save(r, w)
		} else {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		}
	}
	return err
}

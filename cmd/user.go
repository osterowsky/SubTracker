package cmd

import (
	"crypto/md5"
	"encoding/hex"
	d "subtracker/schema/database"
)

func createUser(username, email, password string) (err error) {

	// Generate the MD5 hash of the password
	passwordHash := md5.Sum([]byte(password))

	// Insert the user into the database
	_, err = d.DB.Exec("INSERT INTO users (username, email, passwordhash) VALUES ($1, $2, $3)", username, email, hex.EncodeToString(passwordHash[:]))
	if err != nil {
		return err
	}
	return nil
}

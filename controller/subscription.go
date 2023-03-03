package controller

import (
	"log"
	"net/http"
	"strings"
	d "subtracker/schema"
	"time"
)

func addSub(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")

	startDateRaw, err := time.Parse("2006-01-02", r.FormValue("start-date"))
	if err != nil {
		http.Error(w, "Invalid date format for first payment", http.StatusBadRequest)
		return
	}
	startDate := startDateRaw.Format("2006-01-02")

	billingPeriod := strings.ToLower(r.FormValue("billing-period"))

	// Checks if users chosen any billing period before
	if strings.Contains(billingPeriod, "billing") {
		http.Error(w, "Please choose billing period", http.StatusBadRequest)
		return
	}

	price := r.FormValue("price")

	// get userID from current session
	session, _ := Store.Get(r, "session.id")
	userID := session.Values["user_id"]

	_, err = d.DB.Exec("INSERT INTO subscription (name, start_date, billing_period, price, user_id) VALUES ($1, $2, $3, $4, $5)", name, startDate, billingPeriod, price, userID)
	if err != nil {
		http.Error(w, "Failed to insert subscription data into database", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

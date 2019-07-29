package handlers

import (
	"encoding/json"
	"net/http"
)

// Routes set the routes for the web service

func Routes() {
	http.HandleFunc("/sendjson", SendJSON)
}

// SENDJSON returns a simple JSON document
func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Steven",
		Email: "steven@gmail.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}

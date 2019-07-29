package main

import (
	"log"
	"net/http"

	"./handlers"
)

func main() {
	handlers.Routes()

	log.Println("Listener: Started: Listening on port: 4000")
	http.ListenAndServe(":4000", nil)
}

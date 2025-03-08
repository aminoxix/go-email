package handlers

import (
	"net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	// Write the status code and text to the response
	w.Write([]byte("Connected to server!"))
}

package main

import (
	"log"
	"main/configs"
	"main/handlers"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	// Load environment variables
	configs.LoadEnv()

	bindAddress := ":8080"

	mux := mux.NewRouter()
	prefix := "/api/v0"

	mux.HandleFunc("/", handlers.StatusHandler).Methods("GET")
	mux.HandleFunc(prefix, handlers.StatusHandler).Methods("GET")
	mux.HandleFunc(prefix + "/send-email", handlers.EmailHandler).Methods("POST")

	// Wrap the mux with the CORS middleware
	handlerWithCORS := configs.EnableCORS(mux)

	log.Printf("server is listening at %s", bindAddress)
	log.Fatal(http.ListenAndServe(bindAddress, handlerWithCORS))
}

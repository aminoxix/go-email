package main

import (
	"log"
	"main/handlers"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bindAddress := ":8080"

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.StatusHandler)
	mux.HandleFunc("/send-email", handlers.EmailHandler)

	log.Printf("server is listening at %s", bindAddress)
	log.Fatal(http.ListenAndServe(bindAddress, mux))
}

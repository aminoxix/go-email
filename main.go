package main

import (
	"log"
	"main/configs"
	"main/handlers"
	"net/http"
)

func main() {
	// Load environment variables
	configs.LoadEnv()

	bindAddress := ":8080"

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.StatusHandler)
	mux.HandleFunc("/send-email", handlers.EmailHandler)

	log.Printf("server is listening at %s", bindAddress)
	log.Fatal(http.ListenAndServe(bindAddress, mux))
}

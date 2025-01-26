package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"main/interfaces"
	"main/services"
	"net/http"
	"os"
	"text/template"
)


func EmailHandler(w http.ResponseWriter, r *http.Request) {
	to := os.Getenv("TO_EMAIL")
	from := os.Getenv("FROM_EMAIL")
    password := os.Getenv("GMAIL_APP_PASSWORD")

	// convert to use just first element of slice
	to_slice := []string{to}

	// ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}
	
	// parse JSON request
	var requestBody interfaces.EmailRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	
	subject := "hey, iam " + requestBody.Vars.Name +  " from portfolio, let's chat!"

	// parse HTML template
	tmpl, err := template.ParseFiles("./templates/" + requestBody.Template + ".html")
	if err != nil {
		log.Fatalf("failed to parse template: %v", err)
	}

	// render the template with the map data
	var rendered bytes.Buffer
	if err := tmpl.Execute(&rendered, requestBody.Vars); err != nil {
		log.Fatalf("failed to execute template: %v", err)
	}

	log.Println(rendered.String())

	err = services.SendEmail(
		from,
		password,
		to_slice,
		subject,
		rendered.String(),
	)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("email sent successfully!"))
}

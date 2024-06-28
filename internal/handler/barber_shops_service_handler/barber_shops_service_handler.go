package handlers

import (
	"io"
	"log"
	"net/http"
)

// HandleShopsRequest handles requests for /barber-shop-service/shops endpoint
func HandleShopsRequest(w http.ResponseWriter, r *http.Request) {
	backendURL := "http://localhost:8081/shops"

	// Create a new GET request to the backend service
	req, err := http.NewRequest("GET", backendURL, nil)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Perform the HTTP request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to perform request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		log.Printf("Backend service returned non-200 status code: %d", resp.StatusCode)
		w.WriteHeader(resp.StatusCode)
		return
	}

	// Copy the response from the backend to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := io.Copy(w, resp.Body); err != nil {
		log.Printf("Failed to copy response body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

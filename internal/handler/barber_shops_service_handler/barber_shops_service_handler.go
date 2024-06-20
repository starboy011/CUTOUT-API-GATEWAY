package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// ShopResponse represents the structure of each shop object in the array
type ShopResponse struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	ShopID    string `json:"shopid"`
	ShopName  string `json:"shopname"`
}

// HandleShopsRequest handles requests for /barber-shop-service/shops endpoint
func HandleShopsRequest(w http.ResponseWriter, r *http.Request) {
	// Define the URL of the backend service
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

	// Decode the JSON array response from the backend service
	var responseData []ShopResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		log.Printf("Failed to decode response body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return JSON response to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
}

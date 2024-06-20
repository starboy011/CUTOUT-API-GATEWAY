package handlers

import (
	"encoding/json"
	"net/http"
)

// HandleUsersRequest handles requests for /user-service/users endpoint
func HandleUsersRequest(w http.ResponseWriter, r *http.Request) {
	// Example: Forwarding the request to user service
	// Simulate response from user service
	responseData := map[string]string{
		"message": "Users data from user service",
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
}

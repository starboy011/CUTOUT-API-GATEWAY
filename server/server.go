package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupShopsRoutes sets up nested routes related to shops

// StartServer initializes and starts the HTTP server
func StartServer() error {
	// Initialize mux router
	router := mux.NewRouter()

	// Setup nested routes for shops
	SetupBarberShopsServiceRoutes(router)
	SetupUserServiceRoutes(router)

	// Start HTTP server
	log.Println("Starting API Gateway...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		return err
	}

	return nil
}

package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/starboy011/api-gateway/internal/handler"
)

// StartServer initializes and starts the HTTP server
func StartServer() error {
	// Initialize mux router
	router := mux.NewRouter()

	// Define nested routes for barber-shop-service and user service
	barberShopServiceRouter := router.PathPrefix("/barber-shop-service").Subrouter()
	barberShopServiceRouter.HandleFunc("/shops", handlers.HandleShopsRequest).Methods("GET")

	userServiceRouter := router.PathPrefix("/user-service").Subrouter()
	userServiceRouter.HandleFunc("/users", handlers.HandleUsersRequest).Methods("GET")

	// Start HTTP server
	log.Println("Starting API Gateway...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		return err
	}

	return nil
}

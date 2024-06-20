package server

import (
	"github.com/gorilla/mux"
	handlers "github.com/starboy011/api-gateway/internal/handler/user_service_handler"
)

func SetupUserServiceRoutes(router *mux.Router) {
	// Define nested routes for user service
	userServiceRouter := router.PathPrefix("/user-service").Subrouter()
	userServiceRouter.HandleFunc("/users", handlers.HandleUsersRequest).Methods("GET")

}

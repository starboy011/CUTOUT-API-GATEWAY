package server

import (
	"github.com/gorilla/mux"
	handlers "github.com/starboy011/api-gateway/internal/handler/barber_shops_service_handler"
)

func SetupBarberShopsServiceRoutes(router *mux.Router) {
	// Define nested routes for barber-shop-service
	barberShopServiceRouter := router.PathPrefix("/barber-shop-service").Subrouter()
	barberShopServiceRouter.HandleFunc("/shops", handlers.HandleShopsRequest).Methods("GET")
	barberShopServiceRouter.HandleFunc("/upload-barber-images", handlers.HandleShopsRequest1).Methods("POST")

}

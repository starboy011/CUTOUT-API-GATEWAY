package main

import (
	"fmt"
	"log"

	"github.com/starboy011/api-gateway/server"
)

func main() {

	fmt.Println("Starting server on :8080")
	if err := server.StartServer(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

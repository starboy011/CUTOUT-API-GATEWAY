package main

import (
	"fmt"
	"net/http"

	"github.com/starboy011/CUTOUT-API-GATEWAY/internal/handler"
)

func main() {

	http.HandleFunc("/hello", handler.HelloHandler)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

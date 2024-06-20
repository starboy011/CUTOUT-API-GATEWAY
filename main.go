package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/starboy011/api-gateway/server"
)

func main() {
	// Initialize New Relic agent
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("api-gateway"),
		newrelic.ConfigLicense("eu01xxc12dc4fe3729b337cc4130261eFFFFNRAL"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create a mux router
	router := mux.NewRouter()

	// Setup your routes
	server.SetupBarberShopsServiceRoutes(router)

	// Instrument the handler with New Relic
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		txn := app.StartTransaction(req.URL.Path)
		defer txn.End()

		// Route requests through the mux router
		router.ServeHTTP(w, req)
	})

	// Start your HTTP server
	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

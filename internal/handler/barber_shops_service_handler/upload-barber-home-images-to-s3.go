package handlers

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// HandleShopsRequest1 forwards the POST request to another endpoint
func HandleShopsRequest1(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// Forward the request to the specified endpoint
	forwardedReq, err := http.NewRequest("POST", "http://localhost:8081/upload-barber-shop-home-images", bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "Failed to create forwarded request", http.StatusInternalServerError)
		return
	}

	// Copy the headers from the original request
	forwardedReq.Header = r.Header

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(forwardedReq)
	if err != nil {
		http.Error(w, "Failed to forward request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response from the forwarded request
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	// Set the status code and headers from the forwarded response
	w.WriteHeader(resp.StatusCode)
	for k, v := range resp.Header {
		for _, value := range v {
			w.Header().Add(k, value)
		}
	}

	// Write the response body from the forwarded response
	w.Write(respBody)
}

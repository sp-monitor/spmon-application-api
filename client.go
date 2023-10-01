package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ClientData struct {
	Module    string      `json:"module"`
	Data      interface{} `json:"data"`
	ClientID  string      `json:"client_id"`
	DateTime string      `json:"datetime"`
}

func clientRoutes() {
http.HandleFunc("/client/data", func(w http.ResponseWriter, r *http.Request) {
	// Check that the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON request body into a RequestData struct
	var requestData ClientData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Do something with the request data
	fmt.Printf("Received request with module=%s, client_id=%s, datetime=%s\n", requestData.Module, requestData.ClientID, requestData.DateTime)
	fmt.Printf("Data: %v\n", requestData.Data)

	// Send a response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request received"))

})
}
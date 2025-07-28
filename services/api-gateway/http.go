package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/shared/contracts"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	// validation
	if reqBody.UserID == "" {
		http.Error(w, "user ID is required", http.StatusBadRequest)
		return
	}

	jsonBody, _ := json.Marshal(reqBody)
	reader := bytes.NewReader(jsonBody)
	resp, err := http.Post("http://trip-service:8083/preview", "application/json", reader)
	if err != nil {
		log.Printf("Error calling trip service: %v", err)
		http.Error(w, "failed to call trip service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var respBody any
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := contracts.APIResponse{
		Data:  respBody,
		Error: nil,
	}
	writeJSON(w, http.StatusCreated, response)

}

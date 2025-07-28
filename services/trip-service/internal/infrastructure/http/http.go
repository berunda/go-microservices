package http

import (
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/types"
)

type HttpHandler struct {
	Service domain.TripService
}
type previewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func (h *HttpHandler) HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fare := &domain.RideFareModel{
		UserID: "42",
	}
	ctx := r.Context()
	trip, err := h.Service.CreateTrip(ctx, fare)
	if err != nil {
		log.Println("Error creating trip:", err)
	}
	log.Println("Trip created successfully:", trip)
	writeJSON(w, http.StatusCreated, trip)
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

// In a real application, you would implement the logic to handle the trip preview.

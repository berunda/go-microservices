package main

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"
)

func main() {
	ctx := context.Background()
	inmemRepo := repository.NewInmemRepository()

	fare := &domain.RideFareModel{
		UserID: "42",
	}
	svc := service.NewService(inmemRepo)
	trip, err := svc.CreateTrip(ctx, fare)
	if err != nil {
		log.Println("Error creating trip:", err)
	}
	log.Println("Trip created successfully:", trip)

	// Temp
	for {
		time.Sleep(time.Second)
	}
}

package main

import (
	"log"
	"net/http"

	"paddletraffic/internal/controller"
	"paddletraffic/internal/repository"
	"paddletraffic/internal/service"
)

func main() {
	db, err := repository.OpenDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")

	if err := repository.RunMigrations(db); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize layers
	locationRepo := repository.NewLocationRepository(db)
	locationService := service.NewLocationService(locationRepo)
	locationController := controller.NewLocationController(locationService)

	// Setup routes
	mux := http.NewServeMux()
	locationController.RegisterRoutes(mux)

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

package main

import (
	"log"
	"net/http"

	"github.com/turf-booking-system/internal/config"
)

// @title           Turf Booking System API
// @version         1.0
// @description     A scalable backend for booking box cricket turf slots.
// @host            localhost:8080
// @BasePath        /api/v1

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// TODO: Initialize database connection
	// TODO: Initialize Redis connection
	// TODO: Initialize repositories
	// TODO: Initialize services
	// TODO: Initialize handlers
	// TODO: Setup router with middleware and routes

	log.Printf("server starting on port %s", cfg.ServerPort)

	// Placeholder: basic health check server
	http.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	if err := http.ListenAndServe(":"+cfg.ServerPort, nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/turf-booking-system/internal/config"
	"github.com/turf-booking-system/pkg/cache"
	"github.com/turf-booking-system/pkg/database"
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

	// Root context used for startup connection checks.
	ctx := context.Background()

	// Initialize database connection pool.
	db, err := database.New(ctx, database.Config{
		DSN:             cfg.DBConnectionString(),
		MaxOpenConns:    25,
		MaxIdleConns:    25,
		ConnMaxLifetime: 5 * time.Minute,
	})
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer db.Close()
	log.Println("connected to postgres")

	// Initialize Redis connection.
	rdb, err := cache.New(ctx, cache.Config{
		Addr: cfg.RedisAddr(),
	})
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}
	defer rdb.Close()
	log.Println("connected to redis")

	// TODO: Initialize repositories
	// TODO: Initialize services
	// TODO: Initialize handlers

	// Setup router.
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
	}

	// Configure HTTP server.
	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: router,
	}

	// Start the server in a goroutine so it doesn't block shutdown handling.
	go func() {
		log.Printf("server starting on port %s", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server failed: %v", err)
		}
	}()

	// Wait for an interrupt or terminate signal.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down server...")

	// Give in-flight requests up to 10 seconds to complete.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("server exited cleanly")
}

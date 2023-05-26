package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"vod/internal/api/routes"
	"vod/internal/auth"
	"vod/internal/config"
	"vod/internal/encoding"
	"vod/internal/search"
	"vod/internal/storage"
	"vod/pkg/database"
)

func main() {
	// Load the configuration
	cfg := config.LoadConfig()

	// Initialize database connection
	db := database.NewDatabase(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// Initialize storage service
	storageService := storage.NewStorageService(cfg.VideoStoragePath)

	// Initialize encoding service
	encodingService := encoding.NewEncodingService(cfg.EncodingEnabled, cfg.EncodingFormat, cfg.EncodingPreset)

	// Initialize search service
	searchService := search.NewSearchService()

	// Initialize authentication service
	authService := auth.NewAuthService(cfg.AuthSecret)

	// Initialize the Gin router
	router := gin.Default()

	// Setup the video routes
	videoRoutes := routes.NewVideoRoutes(router, db, storageService, encodingService, searchService, authService)
	videoRoutes.Setup()

	// Start the HTTP server
	serverAddr := fmt.Sprintf(":%d", cfg.ServerPort)
	log.Printf("Server listening on %s", serverAddr)
	err := router.Run(serverAddr)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

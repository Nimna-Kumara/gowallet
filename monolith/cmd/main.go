package main

import (
	"log"

	"github.com/Nimna-Kumara/gowallet/monolith/internal/config"
	"github.com/Nimna-Kumara/gowallet/monolith/internal/database"
	userHandler "github.com/Nimna-Kumara/gowallet/monolith/internal/user/handler"
	userRepository "github.com/Nimna-Kumara/gowallet/monolith/internal/user/repository"
	userService "github.com/Nimna-Kumara/gowallet/monolith/internal/user/service"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting Monolith Wallet Application... ")

	// 1. Load configuration
	cfg := config.LoadConfig()

	// 2. Connet to db with retry
	db, err := database.ConnectWithRetry(cfg.DBDSN)
	if err != nil {
		log.Fatal("Critial Error: Could not connect to database after retries: %w", err)
	}

	defer db.Close()

	log.Println("Application successfully initialized...")

	// initialize layers
	uRepo := userRepository.NewMySQLUserRepository(db)
	uSvc := userService.NewUserService(uRepo)
	uHandler := userHandler.NewUserHandler(uSvc)

	// setup gin router
	router := gin.Default()

	// routes
	router.POST("/api/v1/users", uHandler.Register)
	router.GET("/api/v1/users/:id", uHandler.GetProfile)
	router.PUT("/api/v1/users/:id", uHandler.UpdateProfile)

	// start server
	log.Println("Server running on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed to run: %v", err)
	}
}

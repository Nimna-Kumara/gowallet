package main

import (
	"log"

	"github.com/Nimna-Kumara/gowallet/monolith/internal/config"
	"github.com/Nimna-Kumara/gowallet/monolith/internal/database"
)

func main() {
	log.Println("Starting Monoluth Wallet Application... ")

	// 1. Load configuration
	cfg := config.LoadConfig()

	// 2. Connet to db with retry
	db, err := database.ConnectWithRetry(cfg.DBDSN)
	if err != nil {
		log.Fatal("Critial Error: Could not connect to database after retries: %w", err)
	}

	defer db.Close()

	log.Println("Application successfully initialized...")
}

package main

import (
	"log"
	"net/http"

	"truck-logistics-api/config"
	"truck-logistics-api/db"
	"truck-logistics-api/routes"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to database
	db.ConnectDB()
	defer db.CloseDB() // Ensure database closes on shutdown

	// Initialize routes
	r := routes.SetupRouter()

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

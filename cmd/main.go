package main

import (
	"log"
	"net/http"

	"trucklogistics/config"
	"trucklogistics/db"
	"trucklogistics/routes"
)

func main() {
	// Load config
	config.LoadEnv()

	// Connect to database
	db.ConnectDB()

	// Initialize routes
	r := routes.SetupRouter()

	log.Println("🚀 Server running on port 8080")
	http.ListenAndServe(":8080", r)
}

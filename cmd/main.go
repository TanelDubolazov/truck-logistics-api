package main

import (
	"log"
	"net/http"

	"github.com/TanelDubolazov/truck-logistics-api/db"

	"github.com/TanelDubolazov/truck-logistics-api/routes"

	"github.com/TanelDubolazov/truck-logistics-api/config"
)

func main() {
	// Load config
	config.LoadEnv()

	// Connect to database
	db.ConnectDB()

	// Initialize routes
	r := routes.SetupRouter()

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}

package routes

import (
	"truck-logistics-api/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Truck routes
	r.HandleFunc("/trucks/{id:[0-9]+}", handlers.GetTruckByID).Methods("GET")
	r.HandleFunc("/trucks", handlers.CreateTruck).Methods("POST")

	return r
}

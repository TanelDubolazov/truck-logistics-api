package routes

import (
	"truck-logistics-api/internal/handlers"

	"github.com/gorilla/mux"
)

// SetupRouter initializes routes
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Truck routes
	r.HandleFunc("/trucks/{id:[0-9]+}", handlers.GetTruckByID).Methods("GET")
	r.HandleFunc("/trucks", handlers.CreateTruck).Methods("POST")
	r.HandleFunc("/trucks", handlers.GetAllTrucks).Methods("GET")
	r.HandleFunc("/trucks/{id:[0-9]+}", handlers.UpdateTruck).Methods("PUT")
	r.HandleFunc("/trucks/{id:[0-9]+}", handlers.DeleteTruck).Methods("DELETE")

	return r
}

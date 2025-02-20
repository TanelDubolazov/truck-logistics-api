package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"truck-logistics-api/internal/models"
	"truck-logistics-api/internal/services"

	"github.com/gorilla/mux"
)

// GetAllTrucks fetches all trucks from the database
func GetAllTrucks(w http.ResponseWriter, r *http.Request) {
	trucks, err := services.GetAllTrucks()
	if err != nil {
		http.Error(w, "Failed to fetch trucks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trucks)
}

// GetTruckByID fetches a truck by its ID
func GetTruckByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	truckID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid truck ID", http.StatusBadRequest)
		return
	}

	truck, err := services.GetTruckByID(truckID)
	if err != nil {
		http.Error(w, "Truck not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(truck)
}

// CreateTruck adds a new truck to the database
func CreateTruck(w http.ResponseWriter, r *http.Request) {
	var truck models.Truck
	if err := json.NewDecoder(r.Body).Decode(&truck); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newID, err := services.CreateTruck(&truck)
	if err != nil {
		http.Error(w, "Failed to insert truck", http.StatusInternalServerError)
		return
	}

	truck.ID = newID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(truck)
}

// UpdateTruck modifies an existing truck
func UpdateTruck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	truckID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid truck ID", http.StatusBadRequest)
		return
	}

	var truck models.Truck
	if err := json.NewDecoder(r.Body).Decode(&truck); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = services.UpdateTruck(truckID, &truck)
	if err != nil {
		http.Error(w, "Failed to update truck", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Truck updated successfully"})
}

// DeleteTruck removes a truck from the database
func DeleteTruck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	truckID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid truck ID", http.StatusBadRequest)
		return
	}

	err = services.DeleteTruck(truckID)
	if err != nil {
		http.Error(w, "Failed to delete truck", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Truck deleted successfully"})
}

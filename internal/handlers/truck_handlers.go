package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"truck-logistics-api/db"
	"truck-logistics-api/internal/models"
	"truck-logistics-api/internal/services"

	"github.com/gorilla/mux"
)

func GetAllTrucks(w http.ResponseWriter, r *http.Request) {
	trucks, err := services.GetAllTrucks()
	if err != nil {
		http.Error(w, "Failed to fetch trucks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trucks)
}

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

func CreateTruck(w http.ResponseWriter, r *http.Request) {
	var truck models.Truck
	if err := json.NewDecoder(r.Body).Decode(&truck); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	scheduleJSON, err := json.Marshal(truck.Schedule)
	if err != nil {
		http.Error(w, "Failed to process schedule", http.StatusInternalServerError)
		return
	}

	query := `
		INSERT INTO trucks (load_capacity, ac_status, last_maintenance, expected_maintenance, ac_maintenance, temperature, latitude, longitude, schedule)
		VALUES (:load_capacity, :ac_status, :last_maintenance, :expected_maintenance, :ac_maintenance, :temperature, :latitude, :longitude, :schedule)
		RETURNING id;
	`

	tx := db.DB.MustBegin()
	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		http.Error(w, "Failed to prepare query", http.StatusInternalServerError)
		return
	}

	truckData := map[string]interface{}{
		"load_capacity":        truck.LoadCapacity,
		"ac_status":            truck.ACStatus,
		"last_maintenance":     truck.LastMaintenance,
		"expected_maintenance": truck.ExpectedMaintenance,
		"ac_maintenance":       truck.ACMaintenance,
		"temperature":          truck.Temperature,
		"latitude":             truck.Latitude,
		"longitude":            truck.Longitude,
		"schedule":             string(scheduleJSON),
	}

	var newID int
	err = stmt.Get(&newID, truckData)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to insert truck", http.StatusInternalServerError)
		return
	}
	tx.Commit()

	truck.ID = newID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(truck)
}

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

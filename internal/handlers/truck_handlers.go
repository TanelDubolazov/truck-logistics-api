package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"truck-logistics-api/db"
	"truck-logistics-api/internal/models"

	"github.com/gorilla/mux"
)

// GetTruckByID fetches a truck by its ID
func GetTruckByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	truckID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid truck ID", http.StatusBadRequest)
		return
	}

	var truck models.Truck
	err = db.DB.Get(&truck, "SELECT * FROM trucks WHERE id = $1", truckID)
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

	var newID int
	err = stmt.Get(&newID, truck)
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

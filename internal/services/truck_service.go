package services

import (
	"errors"
	"truck-logistics-api/db"
	"truck-logistics-api/internal/models"
)

func GetTruckByID(truckID int) (*models.Truck, error) {
	var truck models.Truck
	err := db.DB.Get(&truck, "SELECT * FROM trucks WHERE id = $1", truckID)
	if err != nil {
		return nil, errors.New("truck not found")
	}
	return &truck, nil
}

func CreateTruck(truck *models.Truck) (int, error) {
	query := `
		INSERT INTO trucks (load_capacity, ac_status, last_maintenance, expected_maintenance, ac_maintenance, temperature, latitude, longitude, schedule)
		VALUES (:load_capacity, :ac_status, :last_maintenance, :expected_maintenance, :ac_maintenance, :temperature, :latitude, :longitude, :schedule)
		RETURNING id;
	`

	tx := db.DB.MustBegin()
	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		return 0, errors.New("failed to prepare query")
	}

	var newID int
	err = stmt.Get(&newID, truck)
	if err != nil {
		tx.Rollback()
		return 0, errors.New("failed to insert truck")
	}
	tx.Commit()

	return newID, nil
}

func GetAllTrucks() ([]models.Truck, error) {
	var trucks []models.Truck
	err := db.DB.Select(&trucks, "SELECT * FROM trucks")
	if err != nil {
		return nil, errors.New("failed to fetch trucks")
	}
	return trucks, nil
}

func UpdateTruck(truckID int, truck *models.Truck) error {
	query := `
		UPDATE trucks
		SET load_capacity = :load_capacity, ac_status = :ac_status, last_maintenance = :last_maintenance, 
		    expected_maintenance = :expected_maintenance, ac_maintenance = :ac_maintenance, 
		    temperature = :temperature, latitude = :latitude, longitude = :longitude, schedule = :schedule
		WHERE id = :id;
	`
	truck.ID = truckID
	_, err := db.DB.NamedExec(query, truck)
	if err != nil {
		return errors.New("failed to update truck")
	}
	return nil
}

func DeleteTruck(truckID int) error {
	_, err := db.DB.Exec("DELETE FROM trucks WHERE id = $1", truckID)
	if err != nil {
		return errors.New("failed to delete truck")
	}
	return nil
}

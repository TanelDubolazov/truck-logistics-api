package models

import (
	"encoding/json"
	"time"
)

// Truck represents the truck entity
type Truck struct {
	ID                  int             `json:"id"`
	LoadCapacity        float64         `json:"load_capacity"` // Tons
	ACStatus            bool            `json:"ac_status"`
	LastMaintenance     time.Time       `json:"last_maintenance"`
	ExpectedMaintenance time.Time       `json:"expected_maintenance"`
	ACMaintenance       time.Time       `json:"ac_maintenance"`
	Temperature         float64         `json:"temperature"`
	Latitude            float64         `json:"latitude"`
	Longitude           float64         `json:"longitude"`
	Schedule            json.RawMessage `json:"schedule"` // JSON structure
	CreatedAt           time.Time       `json:"created_at"`
}

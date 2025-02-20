package models

import "time"

type Truck struct {
	ID                  int       `db:"id" json:"id"`
	LoadCapacity        float64   `db:"load_capacity" json:"load_capacity"`
	ACStatus            bool      `db:"ac_status" json:"ac_status"`
	LastMaintenance     time.Time `db:"last_maintenance" json:"last_maintenance"`
	ExpectedMaintenance time.Time `db:"expected_maintenance" json:"expected_maintenance"`
	ACMaintenance       time.Time `db:"ac_maintenance" json:"ac_maintenance"`
	Temperature         float64   `db:"temperature" json:"temperature"`
	Latitude            float64   `db:"latitude" json:"latitude"`
	Longitude           float64   `db:"longitude" json:"longitude"`
	Schedule            string    `db:"schedule" json:"schedule"`
	CreatedAt           time.Time `db:"created_at" json:"created_at"`
}

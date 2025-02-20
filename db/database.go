package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectDB() {
	var err error
	connStr := "user=" + os.Getenv("DB_USER") +
		" dbname=" + os.Getenv("DB_NAME") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"

	DB, err = sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	runMigrations()
}

func runMigrations() {
	createTrucksTable := `
	CREATE TABLE IF NOT EXISTS trucks (
		id SERIAL PRIMARY KEY,
		load_capacity DECIMAL(10,2) NOT NULL,
		ac_status BOOLEAN NOT NULL,
		last_maintenance DATE NOT NULL,
		expected_maintenance DATE NOT NULL,
		ac_maintenance DATE NOT NULL,
		temperature DECIMAL(5,2),
		latitude DECIMAL(9,6) NOT NULL,
		longitude DECIMAL(9,6) NOT NULL,
		schedule JSONB NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := DB.Exec(createTrucksTable)
	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	log.Println("Migration completed: Trucks table created/updated.")
}

// CloseDB safely closes the database connection
func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatalf("Error closing DB: %v", err)
	}
	log.Println("Database connection closed.")
}

package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	log.Println("[INFO] Loading .env file...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[ERROR] Error loading .env file")
	}

	log.Println("[INFO] Retrieving DB_URL from environment...")
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("[ERROR] DB_URL is not set in .env file")
	}

	log.Println("[DEBUG] Attempting to open database connection...")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("[ERROR] Failed to open database connection")
	}

	log.Println("[INFO] Successfully connected to the database")

	log.Println("[DEBUG] Starting database migration...")
	if err = migrate(db); err != nil {
		log.Fatal("[ERROR] Database migration failed")
	}

	log.Println("[INFO] Database migrated successfully")

	return db, nil
}

func migrate(db *sql.DB) error {
	log.Println("[INFO] Running database migration queries...")

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS songs (
		"id" SERIAL PRIMARY KEY,
		"group" VARCHAR(255) NOT NULL,
		"song" VARCHAR(255) UNIQUE NOT NULL
	);`

	log.Println("[DEBUG] Executing table creation query...")
	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("[ERROR] Failed to execute request: %v", err)
	}

	return nil
}

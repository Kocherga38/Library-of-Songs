package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not set in .env file")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal()
	}

	log.Println("Successfully connected to the database")

	if err = migrate(db); err != nil {
		log.Fatal("Database migration failed")
	}

	log.Println("Database migrated successfully")

	return db, nil
}

func migrate(db *sql.DB) error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS songs (
		"id" SERIAL PRIMARY KEY,
		"group" VARCHAR(255) NOT NULL,
		"song" VARCHAR(255) UNIQUE NOT NULL
	);`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to execute request: %v", err)
	}

	return nil

}

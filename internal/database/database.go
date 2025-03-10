package database

import (
	"log"
	"os"

	"github.com/Kocherga38/Library-of-Songs/internal/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not set in .env file")
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal()
	}

	log.Println("Successfully connected to the database")

	if err = db.AutoMigrate(&models.Song{}); err != nil {
		log.Fatal("Error migrating database:", err)
	}

	log.Println("Database migrated successfully")

	return db, nil
}

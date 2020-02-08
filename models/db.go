package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

// get environment variable or fallback
func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return fallback
}

// InitDB connects app to postgres databaset
func InitDB() {
	// Get ENV variables for intitializing database
	user := getEnv("POSTGRES_USER", "gorm")
	password := getEnv("POSTGRES_PASSWORD", "gorm")
	database := getEnv("POSTGRES_DB", "gorm")
	host := getEnv("PG_HOST", "localhost")
	port := getEnv("PG_PORT", "5432")

	// Initialize db
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, database, password)
	db, err = gorm.Open("postgres", dbInfo)

	if err != nil {
		log.Println(err)
		panic("Failed to connect to database!")
	}

	log.Println("Setting up the database...")

	// Migrate the schema
	db.AutoMigrate(&Ticket{}, &ScreenShot{}, &FileArtifact{})
}

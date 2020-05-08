package models

import (
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
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
	// Get ENV variables that are always needed
	host := getEnv("PG_HOST", "localhost")
	port := getEnv("PG_PORT", "5432")
	user := getEnv("PG_USER", "gorm")
	password := getEnv("PG_PASSWORD", "gorm")

	// Handle development vs production environment
	var dbConnectionString string
	productionEnv := getEnv("PG_PRODUCTION", "false")
	production, _ := strconv.ParseBool(productionEnv)
	if production {
		// In production, connect to Google Cloud Run's Honeyclient Service
		dbConnectionString = fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, user, password)
	} else {
		// In development, use a local postgres instance
		database := getEnv("PG_DB", "gorm")
		dbConnectionString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, database, password)
	}

	// Connect to database
	db, err = gorm.Open("postgres", dbConnectionString)

	if err != nil {
		log.Println(err)
		panic("Failed to connect to database!")
	}

	log.Println("Setting up the database...")

	// Migrate the schema
	db.AutoMigrate(&Ticket{}, &ScreenShot{}, &FileArtifact{})

	// Get ENV variables for intitializing database
	// password := getEnv("POSTGRES_PASSWORD", "gorm")
	// database := getEnv("POSTGRES_DB", "gorm")
	// dbInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, database, password)

	// projectID := getEnv("PROJECT_ID", "projectID")
	// zone := getEnv("ZONE", "zone")
	// instanceName := getEnv("INSTANCE_NAME", "instanceName")
	// dbName := getEnv("DB_NAME", "dbName")
	// Initialize db
	// var dbConnectionString = fmt.Sprintf("%s@cloudsql(%s:%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC", user, projectID, zone, instanceName, dbName)
	// var dbConnectionString = fmt.Sprintf("%s:%s@cloudsql(%s:%s:%s)/%s", user, password, projectID, zone, instanceName, dbName)
	// dbConnectionString := fmt.Sprintf("%s:%s@unix(/cloudsql/%s:%s:%s)/%s=", user, password, projectID, zone, instanceName, dbName)
	// dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, user, password)
}

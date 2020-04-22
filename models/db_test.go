package models

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
)

var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.Run("postgres", "9.6", []string{"POSTGRES_PASSWORD=testing", "POSTGRES_DB=testing"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	// set expiration to avoid dead instances
	resource.Expire(30)

	if err = pool.Retry(func() error {
		var err error
		testDb, err = sql.Open("postgres", fmt.Sprintf("postgres://postgres:secret@localhost:%s/%s?sslmode=disable", resource.GetPort("5432/tcp"), database))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// When you're done, kill and remove the container
	err = pool.Purge(resource)
}

func TestInitDB(t *testing.T) {
	ConnectToTesting()
}

func ConnectToTesting() {
}

func ResetTables() {
	db.DropTable(&FileArtifact{})
	db.DropTable(&ScreenShot{})
	db.DropTable(&Ticket{})

	db.CreateTable(&FileArtifact{})
	db.CreateTable(&ScreenShot{})
	db.CreateTable(&Ticket{})

	artifacts := []FileArtifact{
		FileArtifact{TicketId: 1, Filename: "test1", Data: []byte("TEST1")},
		FileArtifact{TicketId: 1, Filename: "test2", Data: []byte("TEST2")},
		FileArtifact{TicketId: 2, Filename: "test3", Data: []byte("TEST3")},
		FileArtifact{TicketId: 3, Filename: "test4", Data: []byte("TEST3")},
	}
	for _, a := range artifacts {
		db.Create(&a)
	}

	tickets := []Ticket{
		Ticket{Name: "test1", URL: "test1.com", Processed: false, ScreenShot: nil},
		Ticket{Name: "test2", URL: "test2.com", Processed: false, ScreenShot: nil},
		Ticket{Name: "test3", URL: "test3.com", Processed: false, ScreenShot: nil},
		Ticket{Name: "test4", URL: "test4.com", Processed: false, ScreenShot: nil},
	}
	for _, t := range tickets {
		db.Create(&t)
	}

}

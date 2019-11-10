package models

import (
	"testing"
)

func TestInitDB(t *testing.T) {
	ConnectToTesting()
}

func ConnectToTesting() {
	InitDB("host=127.0.0.1 port=5432 user=testing dbname=testing password=testing sslmode=disable")
	ResetTables()
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

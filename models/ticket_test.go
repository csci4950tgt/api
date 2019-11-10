package models

import (
	"testing"
)

func TestGetTicketById(t *testing.T) {
	ConnectToTesting()
	resp, err := GetTicketById(2)
	if err != nil {
		t.Error("Error when getting ticket by ID")
	}
	if resp.Name != "test2" {
		t.Error("GetTicketById did not return the correct entry")
	}

	resp, err = GetTicketById(10)
	if err == nil {
		t.Error("GetTicketById did not report error when get ticket by an ID which does not exist")
	}
}

func TestGetTickets(t *testing.T) {
	ConnectToTesting()
	tickets, err := GetTickets()
	if err != nil {
		t.Error("Error when get all tickets")
	}
	if len(*tickets) != 4 {
		t.Error("GetTickets did not return all tickets")
	}
}

func TestCreateTicket(t *testing.T) {
	ConnectToTesting()
	ticket := Ticket{Name: "added", URL: "added.com", Processed: true, ScreenShot: nil}
	CreateTicket(&ticket)
	resp, err := GetTicketById(5)
	if err != nil {
		t.Error("Error when getting ticket by ID")
	}
	if resp.Name != "added" {
		t.Error("CreateTicket failed to create a ticket")
	}
}

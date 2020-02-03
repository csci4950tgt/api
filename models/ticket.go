package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

type Ticket struct {
	gorm.Model
	Name       string       `json:"name",gorm:"size:255"`
	URL        string       `json:"url",gorm:"size:4096"`
	Processed  bool         `json:"processed"`
	ScreenShot []ScreenShot `json:"screenshots"`
}

type HoneyclientRequest struct {
	ID          uint         `json: "id"`
	URL         string       `json:"url"`
	ScreenShots []ScreenShot `json:"screenshots"`
}

// ProcessTicket saves processed ticket in database
func ProcessTicket(ticket *Ticket) {
	// Convert ticket body to request format
	req := HoneyclientRequest{ticket.ID, ticket.URL, ticket.ScreenShot}
	if req.ScreenShots == nil {
		// no screenshots in request, set to empty array so honeyclient not mad
		req.ScreenShots = []ScreenShot{}
	}
	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(req)

	// Send POST request to honeyclient to process ticket
	resp, err := http.Post("http://localhost:8000/ticket", "application/json", reqBody)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	// Decode processed ticket for fileArtifact string names
	var body CreateTicketResponse
	json.NewDecoder(resp.Body).Decode(&body)
	if !body.Success {
		log.Println("Error occured in honeyclient while processing ticket.")
	}

	// Loop through file artifact string names, get each file artifact from in memory storage, save in DB
	var fileArtifact FileArtifact
	for _, s := range *body.FileArtifacts {
		// Get file artifact from in memory storage
		resp, err := http.Get("http://localhost:8000" + s)
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()

		// Decode actual file artifact into our struct, set other fields
		json.NewDecoder(resp.Body).Decode(&fileArtifact.Data)
		fileArtifact.TicketId = ticket.ID
		fileArtifact.Filename = s

		// Create file artifact in database
		err = CreateFileArtifact(&fileArtifact)
		if err != nil {
			log.Println(err)
		}
	}

	// Update ticket in db to show done processing
	db.Model(&ticket).Update("processed", true)
	fmt.Println("Honeyclient processed ticket %d", ticket.ID)
}

// Create a ticket in the database
func CreateTicket(ticket *Ticket) error {
	return db.Create(ticket).Error
}

func CreateFileArtifact(fileArtifact *FileArtifact) error {
	return db.Create(fileArtifact).Error
}

func GetTicketById(ID uint) (*Ticket, error) {
	var ticket Ticket
	// Preload line fetches the screenshot table and joins automatically:
	err := db.Preload("ScreenShot", "ticket_id = (?)", ID).First(&ticket, ID).Error

	return &ticket, err
}

func GetTickets() (*[]Ticket, error) {
	var tickets []Ticket
	err := db.Order("created_at DESC").Find(&tickets).Error

	return &tickets, err
}

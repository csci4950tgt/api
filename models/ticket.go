package models

import (
	"bytes"
	"encoding/json"
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

// ProcessTicket saves processed ticket in database
func ProcessTicket(ticket *Ticket) {
	// TODO: implement!!
	requestBody, err := json.Marshal(&ticket)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(ticket)
	resp, err := http.Post("http://localhost:8000/ticket", "application/json", bytes.NewBuffer(requestBody))

	defer resp.Body.Close()

	var body CreateTicketResponse
	json.NewDecoder(resp.Body).Decode(&body)
	log.Println(body)

	if !body.Success {
		log.Fatalln("Error in decoding json")
	}

	var fileArtifact FileArtifact

	for i, s := range *body.FileArtifacts {
		log.Println(i, s)
		resp, err := http.Get("http://localhost:8000" + s)

		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()

		json.NewDecoder(resp.Body).Decode(&fileArtifact.Data)
		fileArtifact.TicketId = ticket.ID
		fileArtifact.Filename = s

		err = CreateFileArtifact(&fileArtifact)
		if err != nil {
			log.Fatalln(err)
		}
	}

	db.Model(&ticket).Update("processed", true)
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

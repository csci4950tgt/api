package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
)

type Ticket struct {
	gorm.Model
	Name           string       `json:"name",gorm:"size:255"`
	URL            string       `json:"url",gorm:"size:4096"`
	Processed      bool         `json:"processed"`
	ScreenShot     []ScreenShot `json:"screenshots"`
	MalwareMatches string       `json:"malwareMatches",gorm:"size:4096"`
}

// ProcessTicket saves processed ticket in database
func ProcessTicket(ticket *Ticket) {
	// Convert ticket body to request format
	if ticket.ScreenShot == nil {
		// no screenshots in request, set to empty array so honeyclient not mad
		ticket.ScreenShot = []ScreenShot{}
	}
	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(ticket)

	// Send POST request to honeyclient to process ticket
	honeyclientStub := os.Getenv("HONEYCLIENT_STUB")
	if honeyclientStub == "" {
		honeyclientStub = "http://localhost:8000"
	}
	resp, err := http.Post(honeyclientStub+"/ticket", "application/json", reqBody)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	// Decode processed ticket for fileArtifact string names
	var body CreateTicketResponse
	json.NewDecoder(resp.Body).Decode(&body)
	if !body.Success {
		log.Println("Error occured in honeyclient while processing ticket.")
		return
	}

	db.Model(&ticket).Update("malwareMatches", body.MalwareMatches)

	// Loop through file artifact string names, get each file artifact from in memory storage, save in DB
	var fileArtifact FileArtifact
	for _, s := range *body.FileArtifacts {
		// Get file artifact from in memory storage
		resp, err := http.Get(honeyclientStub + s)
		if err != nil {
			log.Println(err)
			return
		}
		defer resp.Body.Close()

		// Decode actual file artifact into our struct, set other fields
		resData, err := ioutil.ReadAll(resp.Body)
		fileArtifact.Data = resData
		fileArtifact.TicketId = ticket.ID
		// Remove "/artifacts/<id>" from beginning of 's'
		sArr := strings.Split(s, "/")
		filename := strings.Join(sArr[3:], "/")
		fileArtifact.Filename = filename
		// Create file artifact in database
		err = CreateFileArtifact(&fileArtifact)
		if err != nil {
			log.Println(err)
			return
		}
	}

	// Update ticket in db to show done processing
	defer db.Model(&ticket).Update("processed", true)
	fmt.Println("Honeyclient processed ticket", ticket.ID)
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

package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/csci4950tgt/api/models"
	"github.com/csci4950tgt/api/util"
)

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	// Create a new ticket for handling, encode request into struct
	var ticket models.Ticket

	err := json.NewDecoder(r.Body).Decode(&ticket)

	if err != nil {
		util.WriteHttpErrorCode(w, http.StatusBadRequest, "Object provided is not a valid ticket object.")

		return
	}

	// Create ticket in db
	ticket.Processed = false
	err = models.CreateTicket(&ticket)

	if err != nil {
		util.WriteHttpErrorCode(w, http.StatusInternalServerError, "Failed to create ticket entry for honeyclient to consume. Likely an existing ticket at this ID")

		fmt.Println("Failed to create ticket entry for honeyclient to consume:")
		fmt.Println(err)

		return
	}

	// after returning ticket info to frontend, asynchonously send ticket to honeyclient, save after
	go models.ProcessTicket(&ticket)

	// Initialize Response
	msg := "Successfully created ticket."
	res := models.Response{
		Success: true,
		Message: &msg,
		Ticket:  &ticket,
	}

	util.WriteHttpResponse(w, res)
}

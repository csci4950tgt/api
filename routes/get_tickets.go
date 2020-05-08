package routes

import (
	"fmt"
	"net/http"

	"github.com/csci4950tgt/api/models"
	"github.com/csci4950tgt/api/util"
)

// GetTickets will take a bunch of tickets at the same time.
func GetTickets(w http.ResponseWriter, r *http.Request) {
	// get tickets from database
	tickets, err := models.GetTickets()

	// handle possible error
	if err != nil {
		msg := "Failed to fetch tickets from database."
		util.WriteHttpErrorCode(w, http.StatusInternalServerError, msg)

		fmt.Println(msg)
		fmt.Println(err)

		return
	}

	// Initialize Response
	res := models.Response{
		Success: true,
		Tickets: tickets,
	}

	util.WriteHttpResponse(w, res)
}

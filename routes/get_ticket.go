package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/csci4950tgt/api/models"
	"github.com/csci4950tgt/api/util"
	"github.com/gorilla/mux"
)

// GetTicket will retrive a single ticket from database.
func GetTicket(w http.ResponseWriter, r *http.Request) {
	// Get variables from router
	vars := mux.Vars(r)                       // get dynamic variables from mux handler
	ticketID, err := strconv.Atoi(vars["id"]) // get integer "ID" from var

	if err != nil {
		util.WriteHttpErrorCode(w, http.StatusBadRequest, "Missing required parameter: id.")

		return
	}

	ticket, err := models.GetTicketById(uint(ticketID))

	if err != nil {
		msg := fmt.Sprintf("Failed to find ticket with ID %d.", ticketID)
		util.WriteHttpErrorCode(w, http.StatusNotFound, msg)

		return
	}

	// Initialize Response
	res := models.Response{
		Success: true,
		Ticket:  ticket,
	}

	util.WriteHttpResponse(w, res)
}

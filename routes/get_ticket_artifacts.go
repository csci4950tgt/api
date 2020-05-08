package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/csci4950tgt/api/models"
	"github.com/csci4950tgt/api/util"
	"github.com/gorilla/mux"
)

// GetTicketArtifacts is the function which
func GetTicketArtifacts(w http.ResponseWriter, r *http.Request) {
	// Get variables from route handler
	vars := mux.Vars(r)                       // get dynamic variables from mux handler
	ticketID, err := strconv.Atoi(vars["id"]) // get integer "id" from vars

	if err != nil {
		util.WriteHttpErrorCode(w, http.StatusBadRequest, "Missing required parameter: id.")

		return
	}

	// Make sure the ticket exists:
	_, err = models.GetTicketById(uint(ticketID))

	if err != nil {
		msg := fmt.Sprintf("Failed to find ticket with ID %d.", ticketID)
		util.WriteHttpErrorCode(w, http.StatusNotFound, msg)

		return
	}

	// Fetch the list of matching file artifacts from the database:
	fileArtifacts, err := models.GetAssociatedArtifacts(uint(ticketID))

	if err != nil {
		msg := "Failed to get a list of file artifacts associated with ticket."
		util.WriteHttpErrorCode(w, http.StatusInternalServerError, msg)

		return
	}

	// Initialize Response
	res := models.Response{
		Success:       true,
		FileArtifacts: fileArtifacts,
	}

	util.WriteHttpResponse(w, res)
}

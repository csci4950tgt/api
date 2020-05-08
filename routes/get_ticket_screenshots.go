package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/csci4950tgt/api/models"
	"github.com/csci4950tgt/api/util"
	"github.com/gorilla/mux"
)

// GetTicketScreenshots will retive the screenshots of the request from
// database.
func GetTicketScreenshots(w http.ResponseWriter, r *http.Request) {
	// Get variables from route handler
	vars := mux.Vars(r)                       // get dynamic variables from mux handler
	ticketID, err := strconv.Atoi(vars["id"]) // get integer "id" from vars

	if err != nil {
		util.WriteHttpErrorCode(w, http.StatusBadRequest, "Missing required parameter: id.")

		return
	}

	// Initialize Response
	msg := fmt.Sprintf("Not yet implemented... ticketID: %d", ticketID)
	res := models.Response{
		Success: true,
		Message: &msg,
	}

	util.WriteHttpResponse(w, res)
}

package routes

import (
	"fmt"
	"github.com/csci4950tgt/api/models"
	"github.com/csci4950tgt/api/util"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetHoneyClientById(w http.ResponseWriter, r *http.Request) {
	// If request type is not current request, return error.
	if (*r).Method != "GET" {
		msg := fmt.Sprintf("Method \"%s\" is not allowed.", (*r).Method)
		util.WriteHttpErrorCode(w, http.StatusMethodNotAllowed, msg)

		return
	}

	// Initialize header
	header := http.Header{}
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "GET")

	// Set header
	util.SetHeader(w, header)

	// Initialize Ticket struct
	vars := mux.Vars(r)               // get dynamic variables from mux handler
	id, _ := strconv.Atoi(vars["id"]) // get integer "ID" from var

	ticket, err := models.GetTicketById(uint(id))

	// Initialize Response
	msg := fmt.Sprintf("Fetched ticket '%s' from ID '%d' with URL '%s'", ticket.Name, ticket.ID, ticket.URL)
	res := models.Response{
		Success: true,
		Message: &msg,
	}

	util.WriteHttpResponse(w, res)
}

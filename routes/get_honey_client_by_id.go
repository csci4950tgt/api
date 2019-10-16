package routes

import (
	"fmt"
	"github.com/csci4950tgt/api/models"
	"github.com/csci4950tgt/api/util"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func SetupGetResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methos", "GET")
}

func GetHoneyClientById(w http.ResponseWriter, r *http.Request) {

	// Handling requests.
	SetupGetResponse(&w, r)
	// If request type is not current request, return error.
	if (*r).Method != "GET" {
		util.WriteHttpErrorCode(w, http.StatusMethodNotAllowed, fmt.Sprintf("Method \"%s\" is not allowed.\n", (*r).Method))

		return
	}

	vars := mux.Vars(r)               // get dynamic variables from mux handler
	id, _ := strconv.Atoi(vars["id"]) // get integer "ID" from vars
	ticket := models.Ticket{ID: id, Name: "Hardcoded ticket", URL: "https://example.com"}

	msg := fmt.Sprintf("Fetched ticket '%s' from ID '%d' with URL '%s'", ticket.Name, ticket.ID, ticket.URL)
	res := models.Response{
		Success: true,
		Message: &msg,
	}

	util.WriteHttpResponse(w, res)
}

package routes

import (
	"encoding/json"
	"fmt"
	"github.com/csci4950tgt/api/models"
	"github.com/csci4950tgt/api/util"
	"io/ioutil"
	"net/http"
)

func CreateHoneyClient(w http.ResponseWriter, r *http.Request) {
	// If request type is not current request, return error.
	if (*r).Method != "POST" {
		msg := fmt.Sprintf("Method \"%s\" is not allowed.", (*r).Method)
		util.WriteHttpErrorCode(w, http.StatusMethodNotAllowed, msg)

		return
	}

	// Initialize headers array
	headers := []models.ResponseHeader{
		models.ResponseHeader{
			Key:   "Access-Control-Allow-Origin",
			Value: "*",
		},
		models.ResponseHeader{
			Key:   "Access-Control-Allow-Methods",
			Value: "POST",
		},
	}

	// Set headers
	util.SetHeaders(w, headers)

	// Create a new ticket for handling
	var ticket models.Ticket

	// Decodes json from request into Ticket struct
	json.NewDecoder(r.Body).Decode(&ticket)
	ticket.ID = 1

	// Create and write to json file
	file_name := fmt.Sprintf("../honeyclient/input/%d.json", ticket.ID)
	file, _ := json.MarshalIndent(ticket, "", " ")
	err := ioutil.WriteFile(file_name, file, 0755)

	if err != nil {
		util.WriteHttpErrorCode(w, http.StatusInternalServerError, "Failed to write file for honeyclient to consume.")

		fmt.Println("Failed to write file for honeyclient to consume:")
		fmt.Println(err)

		return
	}

	// TODO: run honeyclient and return error if failed

	// Initialize Response
	msg := fmt.Sprintf("Create ticket '%s' with ID '%d' and URL '%s'", ticket.Name, ticket.ID, ticket.URL)
	res := models.Response{
		Success: true,
		Message: &msg,
	}

	util.WriteHttpResponse(w, res)
}

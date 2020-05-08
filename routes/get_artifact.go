package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/csci4950tgt/api/models"
	"github.com/csci4950tgt/api/util"
	"github.com/gorilla/mux"
)

// GetArtifact which handled by backend from database
func GetArtifact(w http.ResponseWriter, r *http.Request) {
	// Get variables from route handler
	vars := mux.Vars(r)                       // get dynamic variables from mux handler
	ticketID, err := strconv.Atoi(vars["id"]) // get integer "id" from vars

	if err != nil {
		util.WriteHttpErrorCode(w, http.StatusBadRequest, "Missing required parameter: id.")

		return
	}

	fileName := vars["fileName"]

	// Fetch the file artifact:
	fileArtifact, err := models.GetArtifact(uint(ticketID), fileName)

	if err != nil {
		msg := "Failed to get the file artifact."
		util.WriteHttpErrorCode(w, http.StatusInternalServerError, msg)

		fmt.Printf("An error occurred when fetching a file artifact with name %s and ticket ID %d:\n", fileName, ticketID)
		fmt.Println(err)

		return
	}

	// On success, send the file:
	w.WriteHeader(http.StatusOK)
	w.Write(fileArtifact.Data)
}

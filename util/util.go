package util

import (
	"encoding/json"
	"github.com/csci4950tgt/api/models"
	"net/http"
)

// Set all key, value pairs of header in response
func SetHeader(w http.ResponseWriter, header http.Header) {
	for k, _ := range header {
		w.Header().Set(k, header.Get(k))
	}
}

// Encodes a `models.Response` object as JSON, handling errors in writing
// or sending it to the HTTP output stream.
func WriteHttpResponse(w http.ResponseWriter, res models.Response) {
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func WriteHttpError(w http.ResponseWriter, err models.ResponseError) {
	w.WriteHeader(err.Code)

	res := models.Response{
		Success: false,
		Error:   &err,
	}

	WriteHttpResponse(w, res)
}

func WriteHttpErrorCode(w http.ResponseWriter, code int, message string) {
	err := models.ResponseError{
		Code:    code,
		Message: message,
	}

	WriteHttpError(w, err)
}

// Return 200 OK, `{"success": true}`
func WriteHttpEmptySuccess(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)

	res := models.Response{
		Success: true,
	}

	WriteHttpResponse(w, res)
}

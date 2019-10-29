package main

import (
	"fmt"
	"github.com/csci4950tgt/api/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	mux := mux.NewRouter()                                              // create router for handling api endpoints
	mux.HandleFunc("/api/honeyclient/create", routes.CreateHoneyClient) // POST endpoint for creating honeyclient
	mux.HandleFunc("/api/honeyclient/{id}", routes.GetHoneyClientById)  // GET endpoint for getting honeyclient by id
	handler := cors.Default().Handler(mux)
	fmt.Println("Server starting...")
	http.ListenAndServe(":8080", handler) // listen to requests on localhost:8080
}

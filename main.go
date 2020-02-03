package main

import (
	"fmt"
	"net/http"

	"github.com/csci4950tgt/api/models"
	"github.com/csci4950tgt/api/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Handles API routes for mux router
func handleRoutes(r *mux.Router) {
	r.HandleFunc("/api/tickets", routes.GetTickets).Methods("GET")
	r.HandleFunc("/api/tickets", routes.CreateTicket).Methods("POST")
	r.HandleFunc("/api/tickets/{id}", routes.GetTicket).Methods("GET")
	r.HandleFunc("/api/tickets/{id}/artifacts", routes.GetTicketArtifacts).Methods("GET")
	r.HandleFunc("/api/tickets/{id}/artifacts/{fileName:.*}", routes.GetArtifact).Methods("GET")
	r.HandleFunc("/api/tickets/{id}/artifacts/screenshots", routes.GetTicketScreenshots).Methods("GET")
	// TODO: make following route work
	// r.HandleFunc("/api/tickets/{id}/artifacts/js", routes.GetTicketJS).Methods("GET")
}

func main() {
	// Initialize router
	r := mux.NewRouter()
	handleRoutes(r)

	// Initialize DB
	models.InitDB()

	// Listen and serve baby
	fmt.Println("Server starting on http://localhost:8080...")
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000", "http://localhost:5000"})
	http.ListenAndServe(":8080", handlers.CORS(allowedOrigins)(r))
}

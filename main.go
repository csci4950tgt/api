package main

import (
	"fmt"
	"net/http"
	"os"

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
}

func main() {
	// Initialize router
	r := mux.NewRouter()
	handleRoutes(r)

	// Initialize DB
	models.InitDB()

	// Listen and serve baby
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server starting on http://localhost:" + port + "...")
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000", "http://localhost:5000", "https://frontend-bwkgpgz7aq-uc.a.run.app"})
	http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CORS(allowedOrigins)(r))
}

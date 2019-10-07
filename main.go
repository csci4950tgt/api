package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type ScreenShot struct {
	Width    int64  `json:"width"`
	Height   int64  `json:"height"`
	Filename string `json:"filename"`
}

type Ticket struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	URL        string     `json:"url"`
	Screenshot ScreenShot `json:"screenshot"`
}

func main() {
	handler := mux.NewRouter()
	handler.HandleFunc("/api/hello", SayHello) // hello world endpoint
	handler.HandleFunc("/api/createClient", CreateClient)
	fmt.Println("Server starting on localhost:8080...")
	http.ListenAndServe(":8080", handler) // listen to requests on localhost:8080
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!\n")
}

func CreateClient(respond http.ResponseWriter, request *http.Request) {
	var ticket Ticket
	json.NewDecoder(request.Body).Decode(&ticket)
	ticket.ID, ticket.Name = 1, "Dummy ticket"
	fmt.Fprintf(respond, "Ticket %s (%d) will investigate url '%s' with screenshot dimensions %d x %d", ticket.Name, ticket.ID, ticket.URL, ticket.Screenshot.Width, ticket.Screenshot.Height)

}

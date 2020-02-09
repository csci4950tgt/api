package models

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Success       bool            `json:"success"`
	Error         *ResponseError  `json:"error,omitempty"`
	Message       *string         `json:"message,omitempty"`
	Ticket        *Ticket         `json:"ticket,omitempty"`
	Tickets       *[]Ticket       `json:"tickets,omitempty"`
	FileArtifacts *[]FileArtifact `json:"fileArtifacts,omitempty"`
}

type CreateTicketResponse struct {
	Success        bool      `json:"success"`
	FileArtifacts  *[]string `json:"fileArtifacts"`
	MalwareMatches string    `json:"malwareMatches"`
}

// Makes ResponseError satisfy builtin Error type. See: https://blog.golang.org/error-handling-and-go
func (e *ResponseError) Error() string {
	return e.Message
}

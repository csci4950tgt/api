package models

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Success bool           `json:"success"`
	Error   *ResponseError `json:"error,omitempty"`
	Message *string        `json:"message,omitempty"`
}

// Makes ResponseError satisfy builtin Error type. See: https://blog.golang.org/error-handling-and-go
func (e *ResponseError) Error() string {
	return e.Message
}

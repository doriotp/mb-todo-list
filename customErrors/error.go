package customerrors

import "fmt"

type Error struct {
	Code    int
	Message string
}

// Implement the Error method to satisfy the error interface
func (e *Error) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// New creates a new custom error with a given code and message
func New(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

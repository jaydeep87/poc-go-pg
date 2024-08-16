package util

import (
	"fmt"
)

type CustomError struct {
	Code    int
	Message string
}

// Implement the Error() method to satisfy the error interface
func (e CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// Function that returns a custom error
func CreateError(code int, message string) error {
	// Simulate an error condition
	return CustomError{Code: code, Message: message}
}



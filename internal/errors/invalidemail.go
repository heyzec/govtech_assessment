package errors

import (
	"fmt"
	"net/http"
)

type InvalidEmailError struct {
	Query string
}

func NewInvalidEmailError(query string) error {
	return &InvalidEmailError{
		Query: query,
	}
}

func (e *InvalidEmailError) Error() string {
	return fmt.Sprintf("Invalid email: %s", e.Query)
}

func (e *InvalidEmailError) HTTPCode() int {
	return http.StatusBadRequest
}

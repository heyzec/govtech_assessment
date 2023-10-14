package errors

import (
	"fmt"
	"net/http"
)

type MissingFieldError struct {
	Field string
}

func NewMissingFieldError(field string) error {
	return &MissingFieldError{
		Field: field,
	}
}

func (e *MissingFieldError) Error() string {
	return fmt.Sprintf("Missing field: %s", e.Field)
}

func (e *MissingFieldError) HTTPCode() int {
	return http.StatusBadRequest
}

package errors

import (
	"fmt"
	"net/http"
)

type NotFoundError struct {
	ModelName string
	Query     string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found: %s", e.ModelName, e.Query)
}

func (e *NotFoundError) HTTPCode() int {
	return http.StatusNotFound
}

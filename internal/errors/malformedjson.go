package errors

import (
	"fmt"
	"net/http"
)

type MalformedJSONError struct {
	WrappedError error
}

func WrapMalformedJsonError(e error) error {
	return &MalformedJSONError{
		WrappedError: e,
	}
}

func (e *MalformedJSONError) Error() string {
	return fmt.Sprintf("Malformed JSON: %s", e.WrappedError.Error())
}

func (e *MalformedJSONError) HTTPCode() int {
	return http.StatusBadRequest
}

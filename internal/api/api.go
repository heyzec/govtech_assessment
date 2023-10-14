package api

import (
	"encoding/json"
	"net/http"

	"github.com/heyzec/govtech-assignment/internal/errors"
	"gorm.io/gorm"
)

type Response struct {
	Payload json.RawMessage `json:"data,omitempty"`
}

type Handler = func(r *http.Request, db *gorm.DB) (*Response, error)

// Wrapper around http.Handler to abstract use of http.ResponseWriter
func WrapHandler(db *gorm.DB, handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := handler(r, db)

		if err != nil {
			w.Header().Set("Content-Type", "plain/text")

			var i interface{} = err
			exterr, ok := i.(errors.ExternalError)
			if ok {
				w.WriteHeader(exterr.HTTPCode())
			}
			w.Write([]byte(err.Error()))
		} else {
			if response != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(response.Payload)
			} else {
				w.WriteHeader(http.StatusNoContent)
			}
		}

	}
}

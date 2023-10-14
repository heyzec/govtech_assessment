package api

import (
	"encoding/json"
	"net/http"

    "gorm.io/gorm"

)

type Response struct {
    Payload json.RawMessage `json:"data,omitempty"`
    HTTPCode int
}

type Handler = func(r *http.Request, db *gorm.DB) (*Response)

// Wrapper around http.Handler to abstract use of http.ResponseWriter
func WrapHandler(db *gorm.DB, handler Handler) (http.HandlerFunc) {
    return func (w http.ResponseWriter, r *http.Request) {
        response := handler(r, db)
        if response.Payload != nil {
            w.Header().Set("Content-Type", "application/json")
        }
        w.WriteHeader(response.HTTPCode)
        if response.Payload != nil {
            w.Write(response.Payload)
        }
    }
}


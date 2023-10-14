package api

import (
	"log"
	"net/http"

	"github.com/heyzec/govtech-assignment/internal/errors"
	"github.com/heyzec/govtech-assignment/internal/json"
	"gorm.io/gorm"
)

type Response struct {
	//Payload json.RawMessage `json:"data,omitempty"`
	JSONPayload interface{}
}

type Handler = func(r *http.Request, db *gorm.DB) (*Response, error)

// Wrapper around http.Handler to abstract use of http.ResponseWriter
func WrapHandler(db *gorm.DB, handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := handler(r, db)

		if err != nil {
			handleHTTPFailure(w, err)
		} else {
			handleHTTPSuccess(w, res)
		}

	}
}

func handleHTTPSuccess(w http.ResponseWriter, res *Response) {
	if res == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	rawBytes, err := json.EncodeView(res.JSONPayload)
	if err != nil {
		log.Println("Error encoding view")
		return
	}

	w.Write(rawBytes)
}

func handleHTTPFailure(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "plain/text")

	var i interface{} = err
	exterr, ok := i.(errors.ExternalError)
	if ok {
		w.WriteHeader(exterr.HTTPCode())
	}
	w.Write([]byte(err.Error()))
}

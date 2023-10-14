package api

import (
	"log"
	"net/http"

	"github.com/heyzec/govtech-assignment/internal/errors"
	"github.com/heyzec/govtech-assignment/internal/helpers/json"
	"github.com/heyzec/govtech-assignment/internal/views"
	"gorm.io/gorm"
)

type Response struct {
	//Payload json.RawMessage `json:"data,omitempty"`
	JSONPayload interface{}
}

// Wrapper around http.Handler to abstract use of http.ResponseWriter
func WrapHandler[T interface{}](db *gorm.DB,
	handler func(T, *gorm.DB) (*Response, error),
	parser func(*http.Request) (T, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params, err := parser(r)
		if err != nil {
			log.Println("Failed to parse")
		}

		res, err := handler(params, db)

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
	w.Header().Set("Content-Type", "application/json")

	var i interface{} = err
	exterr, ok := i.(errors.ExternalError)
	if ok {
		w.WriteHeader(exterr.HTTPCode())
	}

	view := views.ErrorMessageViewFrom(err.Error())
	rawBytes, err := json.EncodeView(view)
	if err != nil {
		log.Println("Error encoding view")
		return
	}
	w.Write(rawBytes)
}

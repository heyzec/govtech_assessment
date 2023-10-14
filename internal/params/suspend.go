package params

import (
	"log"
	"net/http"

	"github.com/heyzec/govtech-assignment/internal/json"
)

type SuspendParams struct {
	StudentEmail string `json:"student"`
}

func SuspendParamsParseFrom(r *http.Request) (*SuspendParams, error) {
	// Parse request body
	var params SuspendParams
	err := json.DecodeParams(r.Body, &params)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &params, nil
}

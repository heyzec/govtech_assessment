package params

import (
	"net/http"

	"github.com/heyzec/govtech-assignment/internal/errors"
	"github.com/heyzec/govtech-assignment/internal/helpers/emailutils"
	"github.com/heyzec/govtech-assignment/internal/helpers/json"
)

type SuspendParams struct {
	StudentEmail string `json:"student"`
}

func SuspendParamsParseFrom(r *http.Request) (*SuspendParams, error) {
	// Parse request body
	var params SuspendParams
	if err := json.DecodeParams(r.Body, &params); err != nil {
		return nil, errors.WrapMalformedJsonError(err)
	}
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return &params, nil
}

func (params *SuspendParams) Validate() error {
	if params.StudentEmail == "" {
		return errors.NewMissingFieldError("student")
	}
	if err := emailutils.ValidateEmailValid(params.StudentEmail); err != nil {
		return err
	}
	return nil
}

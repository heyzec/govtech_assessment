package params

import (
	"net/http"

	"github.com/heyzec/govtech-assignment/internal/errors"
	"github.com/heyzec/govtech-assignment/internal/helpers/json"
)

type RetrieveForNotificationParams struct {
	TeacherEmail string `json:"teacher"`
	Notification string `json:"notification"`
}

func RetrieveForNotificationParseFrom(r *http.Request) (*RetrieveForNotificationParams, error) {
	// Parse request body
	var params RetrieveForNotificationParams
	if err := json.DecodeParams(r.Body, &params); err != nil {
		return nil, errors.WrapMalformedJsonError(err)
	}
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return &params, nil
}

func (params *RetrieveForNotificationParams) Validate() error {
	if params.TeacherEmail == "" {
		return errors.NewMissingFieldError("teacher")
	}
	return nil
}

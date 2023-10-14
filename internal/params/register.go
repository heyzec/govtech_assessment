package params

import (
	"net/http"

	"github.com/heyzec/govtech-assignment/internal/errors"
	"github.com/heyzec/govtech-assignment/internal/helpers/json"
)

type RegisterStudentsParams struct {
	TeacherEmail  string   `json:"teacher"`
	StudentEmails []string `json:"students"`
}

func RegisterStudentsParseFrom(r *http.Request) (*RegisterStudentsParams, error) {
	// Parse request body
	var params RegisterStudentsParams
	if err := json.DecodeParams(r.Body, &params); err != nil {
		return nil, errors.WrapMalformedJsonError(err)
	}
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return &params, nil
}

func (params *RegisterStudentsParams) Validate() error {
	if params.TeacherEmail == "" {
		return errors.NewMissingFieldError("teacher")
	}
	return nil
}

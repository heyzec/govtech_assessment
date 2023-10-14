package params

import (
	"net/http"

	"github.com/heyzec/govtech-assignment/internal/helpers/emailutils"
)

type CommonStudentsParams struct {
	TeacherEmails []string
}

func CommonStudentsParseFrom(r *http.Request) (*CommonStudentsParams, error) {
	// Parse URL parameter
	teacherEmails := r.URL.Query()["teacher"]
	params := &CommonStudentsParams{
		TeacherEmails: teacherEmails,
	}
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return params, nil
}

func (params *CommonStudentsParams) Validate() error {
	if err := emailutils.ValidateEmailsValid(params.TeacherEmails); err != nil {
		return err
	}
	return nil
}

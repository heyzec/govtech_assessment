package params

import "net/http"

type CommonStudentsParams struct {
	TeacherEmails []string
}

func CommonStudentsParseFrom(r *http.Request) (*CommonStudentsParams, error) {
	// Parse URL parameter
	teacherEmails := r.URL.Query()["teacher"]
	return &CommonStudentsParams{
		TeacherEmails: teacherEmails,
	}, nil
}

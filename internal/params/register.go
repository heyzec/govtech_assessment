package params

import (
	"log"
	"net/http"

	"github.com/heyzec/govtech-assignment/internal/json"
)

type RegisterStudentsParams struct {
	TeacherEmail  string   `json:"teacher"`
	StudentEmails []string `json:"students"`
}

func RegisterStudentsParseFrom(r *http.Request) (*RegisterStudentsParams, error) {
	// Parse request body
	var params RegisterStudentsParams
	err := json.DecodeParams(r.Body, &params)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &params, nil
}

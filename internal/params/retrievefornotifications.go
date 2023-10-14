package params

import (
	"log"
	"net/http"

	"github.com/heyzec/govtech-assignment/internal/helpers/json"
)

type RetrieveForNotificationParams struct {
	TeacherEmail string `json:"teacher"`
	Notification string `json:"notification"`
}

func RetrieveForNotificationParseFrom(r *http.Request) (*RetrieveForNotificationParams, error) {
	// Parse request body
	var params RetrieveForNotificationParams
	err := json.DecodeParams(r.Body, &params)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &params, nil
}

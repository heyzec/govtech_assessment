package routes

import "net/http"

// A simple handler to display friendly message on /
func RootHandler(w http.ResponseWriter, _ *http.Request) {
	friendlyMessage := "Hi there! Please access my APIs via ${HOSTNAME}/api/..."
	w.Write([]byte(friendlyMessage))
}

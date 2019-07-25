package httputil

import (
	"encoding/json"
	"net/http"
)

func WriteJsonResponse(w http.ResponseWriter, status int, body interface{}) error {
	// set content-type
	w.Header().Set("Content-Type", "application/json")
	// set header status code
	w.WriteHeader(status)

	// serialize body
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = w.Write(bodyBytes)
	return err
}

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

func JsonHandler(handler func(*http.Request) (int, interface{})) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// execute request
		status, body := handler(request)

		// override 200 OK with no content to 204 NO CONTENT
		if status == http.StatusOK && body == nil {
			status = http.StatusNoContent
		}

		_ = WriteJsonResponse(writer, status, body)
	}
}

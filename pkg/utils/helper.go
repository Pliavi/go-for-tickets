package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// TODO: Simplify this
func SendJsonResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// TODO: Maybe marshal is not necessary here
	json_str, err := json.Marshal(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong"))
		return
	}

	w.Write([]byte(json_str))
}

// TODO: Use a custom APIError struct instead of error
// SendErrorResponse logs the actual error and sends the error message to the client
func SendErrorResponse(w http.ResponseWriter, statusCode int, message string, err error) {
	log.Fatal(err)
	SendJsonResponse(w, statusCode, map[string]string{"error": message})
}

func SendDefaultErrorResponse(w http.ResponseWriter, err error) {
	SendErrorResponse(w, http.StatusForbidden, "something went wrong", err)
}

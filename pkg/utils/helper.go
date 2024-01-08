package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// TODO: DO NOT USE THIS!!!
func TryPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func SendJsonResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json_str, err := json.Marshal(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong"))
		return
	}

	w.Write([]byte(json_str))
}

func SendErrorResponse(w http.ResponseWriter, statusCode int, message string, err error) {
	log.Fatal(err)
	SendJsonResponse(w, statusCode, map[string]string{"error": message})
}

func SendDefaultErrorResponse(w http.ResponseWriter, err error) {
	SendErrorResponse(w, http.StatusForbidden, "something went wrong", err)
}

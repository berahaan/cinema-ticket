package utils

import (
	"encoding/json"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	response := map[string]string{"message": message}
	json.NewEncoder(w).Encode(response)
}

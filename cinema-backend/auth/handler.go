package auth

import (
	models "cinema-backend/model"
	"encoding/json"
	"io"
	"net/http"
)

// this fucntions main purpose is to act as middleware to protect the protected pages thus front end should call this endpoint to protect pages from middlewares directory ..
func HandleToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	// parse the body as action payload
	var actionPayload models.ActionPayloadAuth
	err = json.Unmarshal(reqBody, &actionPayload)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	result, err := ValidateToken(actionPayload.Input)

	if err != nil {
		errorObject := models.GraphQLError{
			Message: err.Error(),
		}
		errorBody, _ := json.Marshal(errorObject)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorBody)
		return
	}
	data, _ := json.Marshal(result) // this will return a either T or F
	w.Write(data)
}

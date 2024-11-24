package handler

import (
	models "cinema-backend/model"
	"cinema-backend/services"
	"encoding/json"
	"io"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	var actionPayload struct {
		Input models.LoginArgs `json:"input"`
	}

	err = json.Unmarshal(reqBody, &actionPayload)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	result, err := services.Login(actionPayload.Input)
	if err != nil {
		errorObject := models.GraphQLError{
			Message: err.Error(),
		}
		errorBody, _ := json.Marshal(errorObject)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorBody)
		return
	}
	data, _ := json.Marshal(result)
	w.Write(data)
}

package handler

import (
	models "cinema-backend/model"
	"cinema-backend/services"
	"cinema-backend/utils"
	"encoding/json"
	"io"
	"net/http"
)

func Singuphandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorResponse(w, "invalid payload", http.StatusBadRequest)
		return
	}
	var actionPayload models.ActionPayload
	err = json.Unmarshal(reqBody, &actionPayload)
	if err != nil {
		utils.ErrorResponse(w, "invalid payload", http.StatusBadRequest)
		return
	}
	exists, err := utils.GetUserByEmail(actionPayload.Input.Email)

	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(exists.Users) > 0 {
		utils.ErrorResponse(w, "Email already exists please login", http.StatusConflict)
		return
	}
	result, err := services.SignIn(actionPayload.Input)
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	successResponse := models.SuccessResponse{
		ID:      result.ID,
		Message: "account successfully created!",
	}
	data, _ := json.Marshal(successResponse)
	w.Write(data)
}

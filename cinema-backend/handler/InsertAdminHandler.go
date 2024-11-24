package handler

import (
	models "cinema-backend/model"
	"cinema-backend/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func InsertAdminHandler(w http.ResponseWriter, r *http.Request) {
	var adminPayload models.InsertAdminInputPayload
	err := json.NewDecoder(r.Body).Decode(&adminPayload)
	if err != nil {
		log.Printf("Error parsing JSON: %v", err)
		http.Error(w, "Unable to parse JSON", http.StatusBadRequest)
		return
	}
	user, err := services.InsertAdminUser(*adminPayload.Email, adminPayload.Firstname, adminPayload.Lastname, adminPayload.Password, adminPayload.Role)
	if err != nil {
		http.Error(w, "Failed to insert admin user", http.StatusInternalServerError)
		return
	}
	response := models.InsertAdminOutput(user.Insertusersone)
	successResponse := models.InsertAdminOutput{
		ID: response.ID,
	}
	fmt.Fprintln(w, "Admin user inserted successfully")
	data, _ := json.Marshal(successResponse)
	w.Write(data)
}

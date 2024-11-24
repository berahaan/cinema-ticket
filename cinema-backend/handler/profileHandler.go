package handler

import (
	models "cinema-backend/model"
	"cinema-backend/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func HandleProfileUpload(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleProfileUpload called")
	var profileRequest models.ProfileImageInput
	err := json.NewDecoder(r.Body).Decode(&profileRequest)
	if err != nil {
		log.Printf("Error parsing JSON: %v", err)
		utils.ErrorResponse(w, "Unable to parse JSON", http.StatusBadRequest)
		return
	}
	profileImageBase64 := profileRequest.ProfileImage

	if profileImageBase64 == "" {
		log.Println("Profile image is missing in the request payload")
		utils.ErrorResponse(w, "Profile image is required", http.StatusBadRequest)
		return
	}
	fileName := fmt.Sprintf("profile_%d.jpg", time.Now().UnixNano())
	filePath := os.Getenv("PROFILE_FILE_PATH") + fileName
	log.Printf("Saving to FILEPATH: %s", filePath)
	err = utils.SaveBase64Image(profileImageBase64, filePath)
	if err != nil {
		log.Printf("Error during saving to disk: %v", err)
		utils.ErrorResponse(w, "Error saving profile image", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"profileurl": os.Getenv("PROFILE_IMAGE_PATH") + fileName,
		"message":    "profile changed successfully",
	}
	fmt.Println("Saved successfuly")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

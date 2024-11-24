package handler

import (
	"cinema-backend/model"
	"cinema-backend/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleOtherImagesDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Trigger called this to delete other_images")

	w.Header().Set("Content-Type", "application/json")
	var payload models.OtherMovieImagesDeletePayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Debug: Print the entire payload to confirm its structure
	fmt.Printf("Payload decoded successfully: %+v\n", payload)

	// Extract the image URL from the deleted record
	imageURL := payload.Event.Data.Old.OtherImage
	if imageURL == "" {
		fmt.Println("No image URL found in the payload")
		http.Error(w, "Image URL not found", http.StatusBadRequest)
		return
	}
	fmt.Printf("Image URL to delete: %s\n", imageURL)
	filePath, err := utils.ExtractFilePath(imageURL)
	if err != nil {
		fmt.Println("Error while extracting file Paths here", err)
	}
	err = utils.DeleteImageFile(filePath)
	if err != nil {
		log.Printf("Failed to delete image file: %v\n", err)
		http.Error(w, "Failed to delete image file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Other image deleted successfully"))
}

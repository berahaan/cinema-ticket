package handler

import (
	"cinema-backend/model"
	"cinema-backend/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleImageDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here a triggers called this function to delete featured images...  ")
	var payload models.FeatureMovieDeletePayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	imageURL := payload.Event.Data.Old.FeaturedImage
	if imageURL == "" {
		http.Error(w, "Image URL not found", http.StatusBadRequest)
		return
	}
	fmt.Printf("Payload decoded successfully: %+v\n", payload)
	// first we need to extract the filepath here then delete the images from a server
	filePath, err := utils.ExtractFilePath(imageURL)
	if err != nil {
		fmt.Println("Error during extract filepath", err)
		return
	}
	err = utils.DeleteImageFile(filePath)
	if err != nil {
		log.Printf("Failed to delete featured image file here: %v\n", err)
		http.Error(w, "Failed to delete image file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Image deleted successfully"))
}

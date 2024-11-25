package handler

import (
	"cinema-backend/services"
	"encoding/json"
	"fmt"
	"net/http"
)

// UploadFilesHandles handles image uploads (featured and other images).
func UploadFilesHandles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler called for image upload")
	w.Header().Set("Content-Type", "application/json")
	requestData, err := services.ParseImageUploadRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Validate the number of other images
	if len(requestData.OtherImages) > 5 {
		http.Error(w, "Too many other images; max is 5.", http.StatusBadRequest)
		return
	}

	// Initialize response data
	var featuredImageURL string
	var otherImageURLs []string

	// Handle featured image upload
	if len(requestData.FeaturedImage) > 0 {
		fmt.Println("Leng of featured images is 1")
		featuredImageURL, err = services.HandleFeaturedImageUpload(requestData.FeaturedImage)
		if err != nil {
			http.Error(w, "Error saving featured image", http.StatusInternalServerError)
			return
		}
	}
	if len(requestData.OtherImages) > 0 {
		fmt.Println("Length of other images is in handler", len(requestData.OtherImages))
		otherImageURLs, err = services.HandleOtherImagesUpload(requestData.OtherImages)
		if err != nil {
			http.Error(w, "Error saving other images", http.StatusInternalServerError)
			return
		}
	}
	// Build and send the response
	response := services.BuildUploadResponse(featuredImageURL, otherImageURLs)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

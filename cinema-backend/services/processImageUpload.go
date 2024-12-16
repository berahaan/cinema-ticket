package services

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

// this below fucntion help us to just return parsed request data
func ParseImageUploadRequest(r *http.Request) (models.ImageUploadRequest, error) {
	var requestData models.ImageUploadRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		log.Printf("error parsing json: %v", err)
		return requestData, fmt.Errorf("unable to parse json")
	}
	return requestData, nil
}

// this below one is handles featured images here
func HandleFeaturedImageUpload(featuredImageBase64 string) (string, error) {
	featuredImageFilename := fmt.Sprintf("featured_%d.jpg", time.Now().UnixNano())
	featuredImagePath := os.Getenv("FEATURED_FILE_PATH") + featuredImageFilename
	err := utils.SaveBase64Image(featuredImageBase64, featuredImagePath)
	if err != nil {
		log.Printf("Error saving featured image: %v", err)
		return "", err
	}
	return os.Getenv("FEATURED_IMAGE_PATH") + featuredImageFilename, nil
}

// handleOtherImagesUpload saves multiple other images and returns their URLs.
func HandleOtherImagesUpload(otherImagesBase64 []string) ([]string, error) {
	fmt.Println("other images is called here otherImageBase64 length", len(otherImagesBase64))
	var otherImageURLs []string
	for i, imageBase64 := range otherImagesBase64 {
		otherImageFilename := fmt.Sprintf("other_%d_%d.jpg", i+1, time.Now().UnixNano())
		otherImagePath := os.Getenv("OTHERIMG_FILE_PATH") + otherImageFilename
		err := utils.SaveBase64Image(imageBase64, otherImagePath)
		if err != nil {
			log.Printf("Error saving other image %d: %v", i+1, err)
			return nil, err
		}
		url := os.Getenv("OTHER_IMAGE_PATH") + otherImageFilename
		fmt.Println("after saving otherImages ", otherImageFilename)
		otherImageURLs = append(otherImageURLs, url)
	}
	return otherImageURLs, nil
}

// buildUploadResponse creates the JSON response for the upload results.
func BuildUploadResponse(featuredImageURL string, otherImageURLs []string) map[string]interface{} {
	var message string
	if featuredImageURL != "" && len(otherImageURLs) > 0 {
		message = "both images uploaded successfully"
	} else if featuredImageURL != "" {
		message = "Featured image uploaded successfully"
	} else {
		message = "Other images uploaded successfully"
	}
	return map[string]interface{}{
		"featuredurl": featuredImageURL,
		"otherurl":    otherImageURLs,
		"success":     true,
		"message":     message,
	}
}

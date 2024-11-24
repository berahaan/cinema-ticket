package utils

import (
	"fmt"
	"os"
	"strings"
)

// DeleteImageFile deletes the file at the specified filePath.
func DeleteImageFile(filePath string) error {
	fmt.Println("Attempting to delete image:", filePath)
	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}
	fmt.Println("Successfully deleted image:", filePath)
	return nil
}
func ExtractFilePath(url string) (string, error) {
	// Extract the file name from the URL
	fileName := url[strings.LastIndex(url, "/")+1:]
	fmt.Println("Extracted file name:", fileName)
	// Determine the base path dynamically
	if strings.Contains(url, "featuredImage") {
		fmt.Println("Yeah its featured images .....")
		basePath := os.Getenv("FEATURED_FILE_PATH")
		fmt.Println("Base path for featured Images ", basePath)
		if basePath == "" {
			return "", fmt.Errorf("FEATURED_FILE_PATH environment variable is not set")
		}
		// Ensure basePath ends with a "/"
		if !strings.HasSuffix(basePath, "/") {
			basePath += "/"
		}
		return basePath + fileName, nil
	}

	if strings.Contains(url, "otherImages") {
		basePath := os.Getenv("OTHERIMG_FILE_PATH")
		if basePath == "" {
			return "", fmt.Errorf("OTHER_IMAGE_DIRECTORY environment variable is not set")
		}
		// Ensure basePath ends with a "/"
		if !strings.HasSuffix(basePath, "/") {
			basePath += "/"
		}
		return basePath + fileName, nil
	}

	return "", fmt.Errorf("unsupported image type in URL: %s", url)
}

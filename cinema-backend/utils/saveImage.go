package utils

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
)

func SaveBase64Image(data, filePath string) error {
	fmt.Println("FILEPATH HERE ", filePath)
	if strings.HasPrefix(data, "data:image/jpeg;base64,") {
		data = strings.TrimPrefix(data, "data:image/jpeg;base64,")
	} else if strings.HasPrefix(data, "data:image/png;base64,") {
		data = strings.TrimPrefix(data, "data:image/png;base64,")
	} else {
		log.Fatal("Invalid images please choose valid images ")
	}
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Println("Error while decoding here ", err)
		return fmt.Errorf("failed to decode base64 image: %v", err)
	}
	fmt.Printf("Decoded data length: %d\n", len(decodedData))
	if len(decodedData) == 0 {
		fmt.Println("Decoded data have length of  0")
		return fmt.Errorf("failed to decode i think len(decodedData) is 0")
	}
	err = os.WriteFile(filePath, decodedData, 0644)
	if err != nil {
		return fmt.Errorf("failed to save image: %v", err)
	}
	return nil
}

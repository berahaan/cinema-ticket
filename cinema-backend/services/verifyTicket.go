package services

import (
	models "cinema-backend/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func verifyChapaPayment(reference string) (bool, error) {
	apiKey := os.Getenv("CHAPA_SECRET_KEY")
	if apiKey == "" {
		return false, fmt.Errorf("CHAPA_SECRET_KEY is not set")
	}
	url := fmt.Sprintf(os.Getenv("CHAPA_VERIFY_ENDPOINT"), reference)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Authorization", apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return false, fmt.Errorf("unexpected response code %d: %s", resp.StatusCode, string(body))
	}

	var chapaResponse models.ChapaVerificationResponse
	if err := json.NewDecoder(resp.Body).Decode(&chapaResponse); err != nil {
		return false, fmt.Errorf("failed to parse Chapa API response: %w", err)
	}
	fmt.Printf("Chapa verification response: %+v\n", chapaResponse)

	if chapaResponse.Status == "success" && chapaResponse.Data.Status == "success" {
		return true, nil
	}
	return false, nil
}

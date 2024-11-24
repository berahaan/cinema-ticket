package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
)

func VerifyHMAC(secretHash string, body []byte, chapaSignature string) error {
	mac := hmac.New(sha256.New, []byte(secretHash))
	mac.Write(body)
	expectedMAC := mac.Sum(nil)
	decodedChapaSignature, err := hex.DecodeString(chapaSignature)
	if err != nil {
		return fmt.Errorf("invalid Chapa signature format: %w", err)
	}
	if !hmac.Equal(expectedMAC, decodedChapaSignature) {
		return errors.New("HMAC verification failed")
	}
	return nil
}

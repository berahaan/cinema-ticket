package handler

import (
	models "cinema-backend/model"
	"cinema-backend/services"
	"cinema-backend/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	ChapaSecretKey      = "CHAPA_SECRET_KEY"
	WebhookSecretKey    = "WEB_HOOK_SECRET_KEY"
	HasuraAdminSecret   = "HASURA_GRAPHQL_ADMIN_SECRET"
	HasuraEndpoint      = "GRAPHQL_ENDPOINT"
	ChapaVerifyEndpoint = "CHAPA_VERIFY_ENDPOINT"
)

func ChapaWebhookHandler(w http.ResponseWriter, r *http.Request) {
	var payload models.ChapaWebhookPayload
	body, err := io.ReadAll(r.Body)
	chapaSignature := r.Header.Get("x-chapa-signature")
	secretHash := os.Getenv(WebhookSecretKey)

	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &payload); err != nil {
		fmt.Println("Error parsing webhook:", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if chapaSignature == "" || secretHash == "" {
		fmt.Println("Chapa signature or Secret hash is empty")
		return
	}

	// Webhook authentication (HMAC verification)
	if err := utils.VerifyHMAC(secretHash, body, chapaSignature); err != nil {
		fmt.Println("Webhook verification failed:", err)
		http.Error(w, "Invalid webhook signature", http.StatusUnauthorized)
		return
	}

	// Delegate the handling to the payment status processor
	err = services.ProcessPaymentStatus(payload)
	if err != nil {
		fmt.Println("Error processing payment:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook processed successfully"))
}

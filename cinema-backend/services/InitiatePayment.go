package services

import (
	"bytes"
	models "cinema-backend/model"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func GenerateTxRef() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
func InitiatePayment(args models.InitiatePaymentArgs, txRef string) (models.PaymentOutput, error) {
	url := os.Getenv("CHAPA_INITIALIZE_PAYMENT")
	payload := map[string]interface{}{
		"amount":       args.Amount,
		"currency":     args.Currency,
		"email":        args.Email,
		"first_name":   args.FirstName,
		"last_name":    args.LastName,
		"phone_number": args.PhoneNumber,
		"tx_ref":       txRef,
		"callback_url": os.Getenv("CALL_BACK_URL"),
		"return_url":   fmt.Sprintf(os.Getenv("RETURN_URL"), txRef),
		"meta": map[string]interface{}{
			"movie_id":        args.MovieId,
			"schedule_id":     args.ScheduleId,
			"ticket_quantity": args.TicketQuantity,
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return models.PaymentOutput{}, err
	}
	//  fmt.Println("Payload to be sent here",payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return models.PaymentOutput{}, err
	}
	req.Header.Set("Authorization", os.Getenv("CHAPA_SECRET_KEY"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return models.PaymentOutput{}, err
	}
	defer res.Body.Close()

	var chapaResponse models.ChapaResponse
	if err := json.NewDecoder(res.Body).Decode(&chapaResponse); err != nil {
		return models.PaymentOutput{}, err
	}

	if chapaResponse.Status != "success" {
		return models.PaymentOutput{}, fmt.Errorf("chapa API error: %s", chapaResponse.Message)
	}

	return models.PaymentOutput{
		AccessURL: chapaResponse.Data.CheckoutURL,
		Message:   "Sucessfully initiaized payment",
		TextRef:   txRef,
		Status:    200,
	}, nil
}

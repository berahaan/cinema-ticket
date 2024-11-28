package handler

import (
	models "cinema-backend/model"
	"cinema-backend/services"
	// "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler for downloading the ticket
func HandleDownloadTicket(w http.ResponseWriter, r *http.Request) {
	var downloadPayload models.InsertTicketInput
	err := json.NewDecoder(r.Body).Decode(&downloadPayload)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing input payload: %v", err), http.StatusBadRequest)
		return
	}
	// Fetch payment data
	paymentData, err := services.FetchPaymentData(downloadPayload.TxRef)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching payment data: %v", err), http.StatusInternalServerError)
		return
	}

	if len(paymentData.PaymentData) == 0 {
		http.Error(w, "No ticket data found for the provided transaction reference", http.StatusNotFound)
		return
	}
	// Generate the ticket PDF
	// ticketPDF, err := services.GeneratePDFWithQRCode(paymentData.PaymentData[0])
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Error generating ticket PDF: %v", err), http.StatusInternalServerError)
	// 	return
	// }
	// Respond with the PDF file
	response := map[string]string{
		// "pdf_base64": base64.StdEncoding.EncodeToString(ticketPDF),
		"filename":   "ticket.pdf",
	}

	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(response)
	w.Write(data)
}


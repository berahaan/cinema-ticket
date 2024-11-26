package handler

import (
	models "cinema-backend/model"
	"cinema-backend/services"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HandlePaymentInitiate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hasura called this inititae functions now ")
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error here during reading input ", err)
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	// fmt.Println("RespBody::", string(reqBody))
	var actionPayload models.TicketInitializePayload
	if err := json.Unmarshal(reqBody, &actionPayload); err != nil {
		fmt.Println("Error during unmarshal:", err)
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	// Generate tx_ref for transaction
	txRef := services.GenerateTxRef()
	actionPayload.Input.TextRef = txRef

	result, err := services.InitiatePayment(actionPayload.Input, txRef)
	if err != nil {
		fmt.Println("Error during initiate ", err)
		errorObject := models.GraphQLError{Message: err.Error()}
		errorBody, _ := json.Marshal(errorObject)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorBody)
		return
	}

	data, _ := json.Marshal(result)
	w.Write(data)
}

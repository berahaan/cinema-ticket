package services

import (
	"cinema-backend/model"
	"fmt"
)

func ProcessPaymentStatus(payload models.ChapaWebhookPayload) error {
	movieId := payload.Meta.MovieId
	scheduleId := payload.Meta.ScheduleId
	ticketQuantity := payload.Meta.TicketQuantity

	switch payload.Status {
	case "success":
		// Verify payment with Chapa
		verified, err := verifyChapaPayment(payload.TransactionID)
		if err != nil {
			return fmt.Errorf("error verifying payment: %w", err)
		}
		if !verified {
			return fmt.Errorf("payment verification failed")
		}

		if movieId == 0 || scheduleId == 0 || ticketQuantity == 0 {
			return fmt.Errorf("metadata is missing or invalid")
		}

		ticket, err := SendTicketInfoToHasura(movieId, scheduleId, ticketQuantity, payload.FirstName, payload.LastName, payload.TransactionID, payload.Amount, "success", payload.Email, payload.PaymentMethod, payload.Currency, payload.Reference)
		if err != nil {
			return fmt.Errorf("error sending ticket info: %w", err)
		}

		if ticket.InsertTicketOne.TicketId == 0 {
			return fmt.Errorf("failed to insert ticket")
		}
		fmt.Println("Ticket is added successfully")
		return nil

	case "pending":
		fmt.Println("Payment is pending...")
		if movieId == 0 || scheduleId == 0 || ticketQuantity == 0 {
			return fmt.Errorf("metadata is missing or invalid")
		}

		ticket, err := SendTicketInfoToHasura(movieId, scheduleId, ticketQuantity, payload.FirstName, payload.LastName, payload.TransactionID, payload.Amount, "pending", payload.Email, payload.PaymentMethod, payload.Currency, payload.Reference)
		if err != nil {
			return fmt.Errorf("error saving pending payment: %w", err)
		}

		if ticket.InsertTicketOne.TicketId == 0 {
			return fmt.Errorf("failed to insert ticket into database")
		}

		fmt.Println("Ticket added successfully with pending status")

		// Start polling in a goroutine
		go func(transactionID string) {
			verified, err := PollPaymentVerifications(transactionID)
			if err != nil {
				fmt.Printf("Error during polling for transaction %s: %v\n", transactionID, err)
				return
			}
			if verified {
				fmt.Printf("Payment status successfully updated for transaction %s\n", transactionID)
			} else {
				fmt.Printf("Payment verification failed or timed out for transaction %s\n", transactionID)
			}
		}(payload.TransactionID)

		return nil
	case "failed":
		fmt.Println("Payment failed. Updating the status...")
		// Save failed payment status in the database
		if movieId == 0 || scheduleId == 0 || ticketQuantity == 0 {
			return fmt.Errorf("metadata is missing or invalid")
		}

		_, err := SendTicketInfoToHasura(movieId, scheduleId, ticketQuantity, payload.FirstName, payload.LastName, payload.TransactionID, payload.Amount, "failed", payload.Email, payload.PaymentMethod, payload.Currency, payload.Reference)
		if err != nil {
			return fmt.Errorf("error saving failed payment: %w", err)
		}

		fmt.Println("Failed payment status saved successfully")
		return nil

	default:
		fmt.Println("Unexpected payment status:", payload.Status)
		return fmt.Errorf("unexpected payment status: %s", payload.Status)
	}
}

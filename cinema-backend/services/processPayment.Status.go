package services

import (
	models "cinema-backend/model"
	"fmt"
	"log"
	"os"
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
		// first call the fetchPayment data to fetch a userdata here then pass that data to the generate functions her
		paymentData, err := FetchPaymentData(payload.TransactionID)
		if err != nil {
			fmt.Println("Error during the fetching data here in ProcessPayment data ")
			return fmt.Errorf("error while fetching the data %s", err)
		}

		fmt.Println("Now its time to generate the qr code here below ...")
		qrpath, err := GenerateQRCode(paymentData.PaymentData[0])
		if err != nil {
			fmt.Println("Error while sending the data to Qr codes ")
		}
		fmt.Println("QR code generated successfully. Filepath:", qrpath)

		pdfBytes, err := GeneratePDFWithQRCode(paymentData.PaymentData[0], qrpath)
		if err != nil {
			fmt.Println("Error during Generating the pdf..")
			return fmt.Errorf("error while generating pdf")
		}
		savePath := fmt.Sprintf(`./ticket/ticket_%s.pdf`, payload.TransactionID)
		// savePath := "./ticket/ticket.pdf"
		err = os.WriteFile(savePath, pdfBytes, 0644)
		if err != nil {
			log.Fatalf("Failed to save PDF: %v", err)
		}
		fmt.Printf("PDF successfully saved to: %s\n", savePath)
		// recipient, subject, body, attachmentPath
		recepeint := payload.Email
		subject := "Your Movie Ticket"
		body := "Thank you for your purchase! Please find your ticket attached."
		err = SendEmailWithAttachment(recepeint, subject, body, savePath)
		if err != nil {
			fmt.Printf("Error while sending emails here ...")
			return fmt.Errorf("error while sendig emails%s", err)
		}
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
		paymentData, err := FetchPaymentData(payload.TransactionID)
		if err != nil {
			fmt.Println("Error during the fetching data here in ProcessPayment data ")
			return fmt.Errorf("error while fetching the data %s", err)
		}

		fmt.Println("Now its time to generate the qr code here below ...")
		qrpath, err := GenerateQRCode(paymentData.PaymentData[0])
		if err != nil {
			fmt.Println("Error while sending the data to Qr codes ")
		}
		fmt.Println("QR code generated successfully. Filepath:", qrpath)

		pdfBytes, err := GeneratePDFWithQRCode(paymentData.PaymentData[0], qrpath)
		if err != nil {
			fmt.Println("Error during Generating the pdf..")
			return fmt.Errorf("error while generating pdf")
		}
		savePath := fmt.Sprintf(`./ticket/ticket_%s.pdf`, payload.TransactionID)
		// savePath := "./ticket/ticket.pdf"
		err = os.WriteFile(savePath, pdfBytes, 0644)
		if err != nil {
			log.Fatalf("Failed to save PDF: %v", err)
		}
		fmt.Printf("PDF successfully saved to: %s\n", savePath)
		recepeint := payload.Email
		subject := "Your Movie Ticket"
		body := "Thank you for your purchase! Your order is being processsed we will notify you when pending is done thanks ."
		err = SendEmailWithAttachment(recepeint, subject, body, "")
		if err != nil {
			fmt.Printf("Error while sending emails here ...")
			return fmt.Errorf("error while sendig emails%s", err)
		}
		// Start polling in a goroutine
		go func(transactionID string) {
			verified, err := PollPaymentVerifications(transactionID, payload.Email,savePath)
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
		recepeint := payload.Email
		subject := "Your Movie Ticket"
		body := "Thank you for your purchase! But unfortuanltely your order is Failed please try again it later \nThanks "
		err = SendEmailWithAttachment(recepeint, subject, body, "")
		if err != nil {
			fmt.Printf("Error while sending emails here ...")
			return fmt.Errorf("error while sendig emails%s", err)
		}
		return nil

	default:
		fmt.Println("Unexpected payment status:", payload.Status)
		return fmt.Errorf("unexpected payment status: %s", payload.Status)
	}
}

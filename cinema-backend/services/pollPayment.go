package services

import (
	"fmt"
	"time"
)

func PollPaymentVerifications(transactionID string, email string,filepath string) (bool, error) {
	fmt.Println("Poll payment is called now it need to wait for a 15 minutes now")
	maxRetries := 10
	retryInterval := 15 * time.Minute
	attempts := 0
	for attempts < maxRetries {
		fmt.Println("firstly called or polled now for verifications")
		time.Sleep(retryInterval)
		attempts++
		verified, err := verifyChapaPayment(transactionID)
		if err != nil {
			fmt.Printf("Error verifying payment on attempt %d: %v\n", attempts, err)
			continue
		}
		if verified {
			// Update payment status to "paid"
			status := "paid"
			resp, err := UpdatePaymentStatus(transactionID, status)
			if err != nil {
				fmt.Printf("Error updating payment status to 'paid': %v\n", err)
				return false, err
			}
			recepeint := email
			subject := "Your Movie Ticket"
			body := "Thank you for your purchase !\n Your order is successfully done Please find your ticket attached.\n Thanks for Your enjoyment "
			err = SendEmailWithAttachment(recepeint, subject, body,filepath)
			if err != nil {
				fmt.Printf("Error while sending emails here ...")
				return false, fmt.Errorf("error while sendig emails%s", err)
			}

			if resp.UpdateTicketStatus.AffectedRow > 0 {
				fmt.Println("Payment status successfully updated to 'paid'")
				return true, nil
			} else {
				fmt.Println("Failed to update payment status in database")
				return false, fmt.Errorf("no rows affected while updating payment status")
			}
		}
	}
	fmt.Printf("Payment verification timed out after %d attempts for transaction %s\n", maxRetries, transactionID)
	status := "failed"
	_, err := UpdatePaymentStatus(transactionID, status)
	if err != nil {
		fmt.Printf("Error updating payment status to 'failed': %v\n", err)
		return false, err
	}
	       recepeint := email
			subject := "Your Movie Ticket"
			body := "We are sorry that your payment is failed ! Please find your ticket attached."
			err = SendEmailWithAttachment(recepeint, subject, body, "")
			if err != nil {
				fmt.Printf("Error while sending emails here ...")
				return false, fmt.Errorf("error while sendig emails%s", err)
			}
	return false, nil
}

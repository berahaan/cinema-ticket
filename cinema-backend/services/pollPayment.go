package services

import (
	"fmt"
	"time"
)

func PollPaymentVerifications(transactionID string) (bool, error) {
	fmt.Println("Poll payment is called now it need to wait for a 5 minutes now")
	maxRetries := 5
	retryInterval := 10 * time.Second
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
	return false, nil
}

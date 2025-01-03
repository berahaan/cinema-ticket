package services

import (
	"fmt"
	"time"
)

func PollPaymentVerifications(transactionID string, email string, filepath string) (bool, error) {
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
			plainTextBody := `Thank you for your purchase!
			Your order has been successfully completed. Please find your ticket attached.
			We hope you enjoy the movie. Thank you!`

			// HTML version
			htmlBody := `
			<!DOCTYPE html>
			<html>
			<head>
				<style>
					body { font-family: Arial, sans-serif; line-height: 1.6; }
					.container { max-width: 600px; margin: 0 auto; padding: 20px; }
					.header { background-color: #27ae60; color: white; padding: 10px; text-align: center; }
					.content { margin: 20px 0; }
					.footer { font-size: 12px; color: gray; text-align: center; margin-top: 20px; }
				</style>
			</head>
			<body>
				<div class="container">
					<div class="header">
						<h1>Your Movie Ticket</h1>
					</div>
					<div class="content">
						<p>Thank you for your purchase!</p>
						<p>Your order has been successfully completed. Please find your ticket attached.</p>
						<p>We hope you enjoy the movie. Thank you for choosing us!</p>
					</div>
					<div class="footer">
						<p>If you have any questions, contact us at support@yourdomain.com</p>
					</div>
				</div>
			</body>
			</html>
			`
			err = SendEmailWithAttachment(recepeint, subject, plainTextBody, htmlBody, filepath)
			if err != nil {
				fmt.Printf("Error while sending emails here ...")
				return false, fmt.Errorf("error while sending emails: %s", err)
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
	subject := "Your Movie Ticket - Order Failed"
	plainTextBody := "Thank you for your order! Unfortunately, your order has failed. Please try again later. Thank you for your understanding."

	htmlBody := `
<!DOCTYPE html>
<html>
<head>
    <style>
        body { font-family: Arial, sans-serif; line-height: 1.6; }
        .container { max-width: 600px; margin: 0 auto; padding: 20px; }
        .header { background-color: #e74c3c; color: white; padding: 10px; text-align: center; }
        .content { margin: 20px 0; }
        .footer { font-size: 12px; color: gray; text-align: center; margin-top: 20px; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Order Failed</h1>
        </div>
        <div class="content">
            <p>Thank you for your order!</p>
            <p>Unfortunately, your order has failed. Please try again later.</p>
            <p>We apologize for the inconvenience and appreciate your understanding.</p>
        </div>
        <div class="footer">
            <p>If you have any questions, contact us at support@yourdomain.com</p>
        </div>
    </div>
</body>
</html>
`

	// Call the function to send the email
	err = SendEmailWithAttachment(recepeint, subject, plainTextBody, htmlBody, "")
	if err != nil {
		fmt.Printf("Error while sending emails here ...")
		return false, err
	}
	return false, nil
}

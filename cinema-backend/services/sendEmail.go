package services

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func SendEmailWithAttachment(recipient, subject, plainTextBody, htmlBody, attachmentPath string) error {
	m := gomail.NewMessage()

	// Set email headers
	m.SetHeader("From", os.Getenv("EMAIL"))
	m.SetHeader("To", "berahaan@gmail.com")
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", plainTextBody)
	if htmlBody != "" {
		m.AddAlternative("text/html", htmlBody)
	}
	password := os.Getenv("GOOGLE_PASSWORD")
	emailme := os.Getenv("EMAIL")
	if password == "" || emailme == "" {
		return fmt.Errorf("error both email and password empty are empty")
	}
	// Attach the file if provided
	if attachmentPath == "" {
		fmt.Println("Attachment is missing")
	} else {
		m.Attach(attachmentPath)
	}

	// Configure the SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "berahaan@gmail.com", "ciak utkb onhw ngll")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("could not send email: %w", err)
	}

	fmt.Println("Email sent successfully!")
	return nil
}

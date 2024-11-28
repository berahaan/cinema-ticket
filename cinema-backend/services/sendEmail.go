package services

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendEmailWithAttachment(recipient, subject, body, attachmentPath string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", "berahaan@gmail.com")
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)
	if attachmentPath == "" {
		fmt.Println("Attachment is missing")
	}
	if attachmentPath != "" {
		m.Attach(attachmentPath)
	}
	d := gomail.NewDialer("smtp.gmail.com", 587, "berahaan@gmail.com", "puho wido fknz bilb")
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("could not send email: %w", err)
	}
	fmt.Println("Email sent successfully!")
	return nil
}

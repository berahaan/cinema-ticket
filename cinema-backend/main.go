package main

import (
	// "cinema-backend/services"
	// models "cinema-backend/model"
	"cinema-backend/auth"
	"cinema-backend/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

// email := "cinemaManage@gmail.com"
// firstName := "Ifaa"
// lastName := "Habtamu"
// password := "manager_password"
// role := "managers"
// user, err := services.InsertAdminUser(email, firstName, lastName, password, role)
// if err != nil {
// 	log.Fatal("Failed to insert admin user", http.StatusInternalServerError)
// 	return
// }
// response := models.InsertAdminOutput(user.Insertusersone)
// if response.ID > 0 {
// 	fmt.Println("Admin user successfully added ")
// }

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error", err)
	}
	// recipient := "berahaan@gmail.com"
	// subject := "Your Movie Ticket"
	// body := "Thank you for your purchase! Please find your ticket attached."
	// attachmentPath := "./ticket/output_ticket.pdf"
	// err = services.SendEmailWithAttachment(recipient, subject, body, attachmentPath)
	// if err != nil {
	// 	fmt.Println("Failed to send email:", err)
	// } else {
	// 	fmt.Println("Email sent successfully!")
	// }

	router := routes.SetupRoutes()
	log.Println("Server running on port 4000")
	err = http.ListenAndServe(":4000", auth.EnableCORS(router))
	if err != nil {
		log.Fatal("error ", err)
	}
}

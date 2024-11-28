package services

import (
	models "cinema-backend/model"
	"fmt"
	"os"
	"github.com/skip2/go-qrcode"
)

// generateQRCode generates a QR code from the given ticket data and saves it as an image.
func GenerateQRCode(ticketData models.TicketData) (string, error) {
	// Define the path for the QR code image
	folderPath := "./qrcodes"
    // Create the folder if it doesn't exist
    err := os.MkdirAll(folderPath, os.ModePerm)
    if err != nil {
        return "", fmt.Errorf("failed to create folder: %v", err)
    }
    filename := fmt.Sprintf("%s_qr.png", ticketData.TxRef)
    filepath := fmt.Sprintf("%s/%s", folderPath, filename)
	ticket := fmt.Sprintf(
		"tx_ref: %s\nMovie Title: %s\nDate: %s\nStart Time: %s\nEnd Time: %s\nCinema Hall: %s\nTickets: %d\nName: %s %s\nEmail: %s\nAmount: %s",
		string(ticketData.TxRef),
		string(ticketData.Movie.Title),
		string(ticketData.MovieSchedule.ScheduleDate),
		string(ticketData.MovieSchedule.StartTime),
		string(ticketData.MovieSchedule.EndTime),
		string(ticketData.MovieSchedule.CinemaHall),
		int(ticketData.TicketQuantity), 
		string(ticketData.Firstname),
		string(ticketData.Lastname),
		string(ticketData.Email),
		string(ticketData.Amount),
	)
	// Generate the QR code and save it as an image file
	err= qrcode.WriteFile(ticket, qrcode.Medium, 256, filepath)
	if err != nil {
		return "", fmt.Errorf("could not generate QR code: %v", err)
	}

	return filepath, nil
}


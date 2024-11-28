package services
import (
	"time"
	"fmt"
	"bytes"
	"cinema-backend/model"
	"github.com/jung-kurt/gofpdf"
)

func GeneratePDFWithQRCode(ticket models.TicketData, qrCodePath string) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "", "")
	pdf.SetAutoPageBreak(false, 0)
	pdf.AddPageFormat("P", gofpdf.SizeType{Wd: 180, Ht: 120})
	pdf.SetDrawColor(0, 0, 0)
	pdf.SetLineWidth(1.5)
	pdf.RoundedRect(5, 5, 170, 110, 5, "1234", "D")
	pdf.SetFillColor(245, 245, 245) 
	pdf.Rect(5, 5, 170, 110, "F")   

	pdf.SetFont("Arial", "B", 18)
	pdf.SetTextColor(255, 87, 34) 
	pdf.CellFormat(170, 12, "ETHIO CINEMA TICKET", "", 1, "C", false, 0, "")
	pdf.Ln(5)

	// Movie title
	pdf.SetFont("Arial", "B", 14)
	pdf.SetTextColor(0, 0, 0) // Black
	pdf.CellFormat(170, 8, fmt.Sprintf("Movie Title:: %s", ticket.Movie.Title), "", 1, "C", false, 0, "")
	pdf.Ln(5)

	scheduleDate, err := time.Parse("2006-01-02", string(ticket.MovieSchedule.ScheduleDate))
	if err != nil {
		return nil, fmt.Errorf("failed to parse schedule date: %w", err)
	}
	weekday := scheduleDate.Weekday().String() 
	formattedDate := fmt.Sprintf("%s, %s", weekday, scheduleDate.Format("01/02/2006"))

	// Schedules section
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(170, 6, "Schedule Details:", "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "", 11)
	pdf.CellFormat(170, 6, fmt.Sprintf("Scheduled Date: %s", formattedDate), "", 1, "C", false, 0, "")
	pdf.Ln(2) 
	pdf.CellFormat(170, 6, fmt.Sprintf("Time: %s - %s",
		formatTime(string(ticket.MovieSchedule.StartTime)),
		formatTime(string(ticket.MovieSchedule.EndTime)),
	), "", 1, "C", false, 0, "")
	pdf.Ln(5) 
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(170, 6, "Customer Information:", "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "", 11)
	pdf.CellFormat(170, 6, fmt.Sprintf("Name: %s %s", ticket.Firstname, ticket.Lastname), "", 1, "C", false, 0, "")
	pdf.CellFormat(170, 6, fmt.Sprintf("Email: %s", ticket.Email), "", 1, "C", false, 0, "")
	pdf.Ln(5)
	// Payment Details
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(170, 6, "Payment Details:", "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "", 11)
	pdf.CellFormat(170, 6, fmt.Sprintf("Number of Tickets: %d", ticket.TicketQuantity), "", 1, "C", false, 0, "")
	pdf.CellFormat(170, 6, fmt.Sprintf("Amount Paid(Birr): %s", ticket.Amount), "", 1, "C", false, 0, "")
	pdf.Ln(5)
	if qrCodePath != "" {
		xPos := 130.0 
		yPos := 70.0
		width := 30.0 
		height := 30.0 
		pdf.Image(qrCodePath, xPos, yPos, width, height, false, "", 0, "")
	}

	// Footer
	pdf.SetFont("Arial", "I", 10)
	pdf.SetTextColor(50, 50, 50)
	pdf.SetXY(5, 110)
	pdf.CellFormat(170, 6, "Thank you for your purchase! Enjoy the movie!", "", 1, "C", false, 0, "")

	var buffer bytes.Buffer
	err = pdf.Output(&buffer)
	if err != nil {
		return nil, fmt.Errorf("failed to generate PDF: %w", err)
	}

	return buffer.Bytes(), nil
}
// Helper function to format time
func formatTime(input string) string {
	t, err := time.Parse("2006-01-02T15:04:05", input)
	if err != nil {
		return input
	}
	return t.Format("3:04 PM")
}

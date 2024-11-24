package handler

import (
	"bytes"
	"cinema-backend/utils"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/diegosz/go-graphql-client"
	"github.com/jung-kurt/gofpdf"
)

// Input payload from frontend
type InsertTicketInput struct {
	TxRef string `json:"tx_ref"`
}

// Struct for fetching payment data
type TicketData struct {
	Firstname      graphql.String `graphql:"firstname"`
	Lastname       graphql.String `graphql:"lastname"`
	Email          graphql.String `graphql:"email"`
	TxRef          graphql.String `graphql:"tx_ref"`
	Amount         graphql.String `graphql:"amount"`
	TicketQuantity graphql.Int    `graphql:"ticket_quantity"`
	PaymentMethod  graphql.String `graphql:"payment_method"`
	PaymentStatus  graphql.String `graphql:"payment_status"`
	Movie          struct {
		Title    graphql.String `graphql:"title"`
		Director struct {
			Name graphql.String `graphql:"name"`
		} `graphql:"director"`
	} `graphql:"movie"`
	MovieSchedule struct {
		CinemaHall   graphql.String `graphql:"cinema_hall"`
		StartTime    graphql.String `graphql:"start_time"`
		EndTime      graphql.String `graphql:"end_time"`
		ScheduleDate graphql.String `graphql:"schedule_date"`
	} `graphql:"movie_schedule"`
}

// Output structure for fetched data
type TicketFetchOutput struct {
	PaymentData []TicketData `json:"payment_data"`
}

// Handler for downloading the ticket
func HandleDownloadTicket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleDownloadTicket function called")

	// Decode the input payload
	var downloadPayload InsertTicketInput
	err := json.NewDecoder(r.Body).Decode(&downloadPayload)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing input payload: %v", err), http.StatusBadRequest)
		return
	}

	fmt.Println("Input payload from frontend:", downloadPayload)

	// Fetch payment data
	paymentData, err := fetchPaymentData(downloadPayload.TxRef)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching payment data: %v", err), http.StatusInternalServerError)
		return
	}

	if len(paymentData.PaymentData) == 0 {
		http.Error(w, "No ticket data found for the provided transaction reference", http.StatusNotFound)
		return
	}

	// Generate the ticket PDF
	ticketPDF, err := generatePDF(paymentData.PaymentData[0])
	if err != nil {
		http.Error(w, fmt.Sprintf("Error generating ticket PDF: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the PDF file
	response := map[string]string{
		"pdf_base64": base64.StdEncoding.EncodeToString(ticketPDF),
		"filename":   "ticket.pdf",
	}

	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(response)
	w.Write(data)
}

// Function to fetch payment data from Hasura using GraphQL
func fetchPaymentData(tx_ref string) (TicketFetchOutput, error) {
	// Prepare a query to fetch data
	var query struct {
		PaymentData []TicketData `graphql:"ticket(where: {tx_ref: {_eq: $tx_ref}})"`
	}

	httpClient := &http.Client{
		Transport: &utils.CustomTransport{
			Transport: http.DefaultTransport,
			Headers: map[string]string{
				"x-hasura-admin-secret": os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET"),
			},
		},
	}
	client := graphql.NewClient(os.Getenv("GRAPHQL_ENDPOINT"), httpClient)

	// Define the query variables
	variables := map[string]interface{}{
		"tx_ref": graphql.String(tx_ref),
	}

	fmt.Println("GraphQL Query Variables:", variables)

	// Execute the query
	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return TicketFetchOutput{}, fmt.Errorf("failed to execute query: %w", err)
	}

	fmt.Println("Query result:", query)
	return TicketFetchOutput{
		PaymentData: query.PaymentData,
	}, nil
}

func generatePDF(ticket TicketData) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "", "") // Custom dimensions
	pdf.SetAutoPageBreak(false, 0)
	pdf.AddPageFormat("P", gofpdf.SizeType{Wd: 180, Ht: 120}) // Adjusted ticket dimensions

	// Full-page border
	pdf.SetDrawColor(0, 0, 0)
	pdf.SetLineWidth(1.5)
	pdf.RoundedRect(5, 5, 170, 110, 5, "1234", "D")

	// Background color inside the border
	pdf.SetFillColor(245, 245, 245) // Light gray
	pdf.Rect(5, 5, 170, 110, "F")   // Filled rectangle inside the border

	// Title
	pdf.SetFont("Arial", "B", 18)
	pdf.SetTextColor(255, 87, 34) // Orange
	pdf.CellFormat(170, 12, "ETHIO CINEMA TICKET", "", 1, "C", false, 0, "")
	pdf.Ln(5)

	// Movie title
	pdf.SetFont("Arial", "B", 14)
	pdf.SetTextColor(0, 0, 0) // Black
	pdf.CellFormat(170, 8, fmt.Sprintf("Movie Title:: %s", ticket.Movie.Title), "", 1, "C", false, 0, "")
	pdf.Ln(5)

	// Format the schedule date to "Wednesday, 11/20/2024"
	scheduleDate, err := time.Parse("2006-01-02", string(ticket.MovieSchedule.ScheduleDate))
	if err != nil {
		return nil, fmt.Errorf("failed to parse schedule date: %w", err)
	}
	weekday := scheduleDate.Weekday().String() // Get the weekday
	formattedDate := fmt.Sprintf("%s, %s", weekday, scheduleDate.Format("01/02/2006"))

	// Schedules section
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(170, 6, "Schedule Details:", "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "", 11)
	pdf.CellFormat(170, 6, fmt.Sprintf("Scheduled Date: %s", formattedDate), "", 1, "C", false, 0, "")
	pdf.Ln(2) // Added space between the date and time

	// Time
	pdf.CellFormat(170, 6, fmt.Sprintf("Time: %s - %s",
		formatTime(string(ticket.MovieSchedule.StartTime)),
		formatTime(string(ticket.MovieSchedule.EndTime)),
	), "", 1, "C", false, 0, "")
	pdf.Ln(5) // Added space before customer info

	// Customer information
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

	// Return the PDF as a byte slice
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

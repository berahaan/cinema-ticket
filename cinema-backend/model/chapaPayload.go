package models

type ChapaWebhookPayload struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"mobile"`
	Status        string `json:"status"`
	Amount        string `json:"amount"`
	Reference     string `json:"reference"`
	CreatedAt     string `json:"created_at"`
	PaymentMethod string `json:"payment_method"`
	Currency      string `json:"currency"`
	TransactionID string `json:"tx_ref"`
	Meta          struct {
		MovieId        int `json:"movie_id"`
		ScheduleId     int `json:"schedule_id"`
		TicketQuantity int `json:"ticket_quantity"`
	} `json:"meta"`
}

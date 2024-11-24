package models

type Action struct {
	Name string `json:"name"`
}
type GraphQLError struct {
	Message string `json:"message"`
}
type SessionVariables struct {
	Role   string `json:"x-hasura-role"`
	UserID string `json:"x-hasura-user-id"`
}
type InitiatePaymentArgs struct {
	MovieId        int    `json:"movie_id"`
	ScheduleId     int    `json:"schedule_id"`
	TicketQuantity int    `json:"ticket_quantity"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Currency       string `json:"currency"`
	Email          string `json:"email"`
	Amount         int    `json:"amount"`
	ReturnURL      string `json:"return_url"`
	PhoneNumber    string `json:"phone_number"`
	TextRef        string `json:"tx_ref"`
}
type TicketInitializePayload struct {
	Action           Action              `json:"action"`
	Input            InitiatePaymentArgs `json:"input"`
	RequestQuery     string              `json:"request_query"`
	SessionVariables SessionVariables    `json:"session_variables"`
}
type PaymentOutput struct {
	AccessURL string `json:"access_url"`
	Message   string `json:"message"`
	Status    int    `json:"status"`
	TextRef   string `json:"tx_ref"`
}
type ChapaResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		CheckoutURL string `json:"checkout_url"`
	} `json:"data"`
}
type ChapaVerificationResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Amount    float64 `json:"amount"`
		Currency  string  `json:"currency"`
		Status    string  `json:"status"`
		Reference string  `json:"tx_ref"`
	} `json:"data"`
}
type UpdateResponseStatus struct {
	UpdateTicketStatus struct {
		AffectedRow int `graphql:"affected_row"`
	} `graphql:"update_ticket"`
}
type InsertPaymentOutput struct {
	InsertTicketOne struct {
		TicketId int `graphql:"ticket_id"`
	} `graphql:"insert_ticket_one"`
}

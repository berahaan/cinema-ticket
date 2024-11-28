package models
import (
	"github.com/diegosz/go-graphql-client"
)
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
type TicketFetchOutput struct {
	PaymentData []TicketData `json:"payment_data"`
}
type InsertTicketInput struct {
	TxRef string `json:"tx_ref"`
}
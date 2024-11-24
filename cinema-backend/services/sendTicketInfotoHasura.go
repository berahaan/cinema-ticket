package services

import (
	models "cinema-backend/model"
	"cinema-backend/utils"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/diegosz/go-graphql-client"
)

func SendTicketInfoToHasura(MovieId int, ScheduleId int, TicketQuantity int, firstName, lastName, txRef, amount, status, email, paymentMethod, currency, reference string) (response models.InsertPaymentOutput, err error) {
	if status == "success" {
		status = "paid"
	}
	hasuraEndpoint := os.Getenv("GRAPHQL_ENDPOINT")
	adminSecret := os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET")
	if hasuraEndpoint == "" || adminSecret == "" {
		return response, fmt.Errorf("hasura graphql endpoint environment variable is not set or admins secret is not set")
	}
	httpClient := &http.Client{
		Transport: &utils.CustomTransport{
			Transport: http.DefaultTransport,
			Headers: map[string]string{
				"x-hasura-admin-secret": adminSecret,
			},
		},
	}
	client := graphql.NewClient(hasuraEndpoint, httpClient)
	var mutation struct {
		InsertTicketOne struct {
			TicketID int `graphql:"ticket_id"`
		} `graphql:"insert_ticket_one(object:{movie_id:$MovieId,ticket_quantity: $TicketQuantity,schedule_id: $ScheduleId,firstname: $firstName,lastname: $lastName,tx_ref: $txRef,amount: $amount,payment_status: $status,email: $email,payment_method: $paymentMethod,currency: $currency,payment_reference: $reference})"`
	}
	variables := map[string]interface{}{
		"TicketQuantity": graphql.Int(TicketQuantity),
		"MovieId":        graphql.Int(MovieId),
		"ScheduleId":     graphql.Int(ScheduleId),
		"firstName":      graphql.String(firstName),
		"lastName":       graphql.String(lastName),
		"txRef":          graphql.String(txRef),
		"amount":         graphql.String(amount),
		"status":         graphql.String(status),
		"email":          graphql.String(email),
		"paymentMethod":  graphql.String(paymentMethod),
		"currency":       graphql.String(currency),
		"reference":      graphql.String(reference),
	}
	err = client.Mutate(context.Background(), &mutation, variables)
	if err != nil {
		fmt.Println("Unable to execute mutations ")
		log.Printf("GraphQL mutation error: %v\n", err)
		return response, fmt.Errorf("failed to execute mutation: %w", err)
	}
	response.InsertTicketOne.TicketId = mutation.InsertTicketOne.TicketID
	fmt.Println("Response from Mutations:::", response.InsertTicketOne.TicketId)
	return response, nil
}

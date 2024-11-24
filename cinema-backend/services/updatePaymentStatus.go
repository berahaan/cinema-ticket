package services

import (
	models "cinema-backend/model"
	"cinema-backend/utils"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/diegosz/go-graphql-client"
)

func UpdatePaymentStatus(txref string, status string) (response models.UpdateResponseStatus, err error) {
	// Fetch Hasura endpoint and admin secret
	hasuraEndpoint := os.Getenv("GRAPHQL_ENDPOINT")
	hasuraAdminSecret := os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET")
	if hasuraEndpoint == "" || hasuraAdminSecret == "" {
		return response, fmt.Errorf("HASURA_GRAPHQL_ENDPOINT or HASURA_ADMIN_SECRET environment variable is not set")
	}

	fmt.Printf("Updating payment status for TxRef: %s with status: %s\n", txref, status)

	// Create a custom HTTP client with necessary headers
	httpClient := &http.Client{
		Timeout: 15 * time.Second,
		Transport: &utils.CustomTransport{
			Transport: http.DefaultTransport,
			Headers: map[string]string{
				"x-hasura-admin-secret": hasuraAdminSecret,
			},
		},
	}

	// Initialize GraphQL client
	client := graphql.NewClient(hasuraEndpoint, httpClient)

	// Define GraphQL mutation
	var mutation struct {
		UpdatePayment struct {
			AffectedRow int `graphql:"affected_rows"`
		} `graphql:"update_ticket(_set:{payment_status:$status},where: {tx_ref: {_eq: $tx_ref}})"`
	}

	// Set variables for the mutation
	variables := map[string]interface{}{
		"status": graphql.String(status),
		"tx_ref": graphql.String(txref),
	}

	// Execute the mutation
	err = client.Mutate(context.Background(), &mutation, variables)
	if err != nil {
		log.Printf("GraphQL mutation error for TxRef: %s: %v\n", txref, err)
		return response, fmt.Errorf("failed to execute mutation for TxRef %s: %w", txref, err)
	}
	// Check the result of the mutation
	if mutation.UpdatePayment.AffectedRow == 0 {
		log.Printf("No rows were updated for TxRef: %s. It might not exist in the database.\n", txref)
		return response, fmt.Errorf("no rows were updated for TxRef: %s", txref)
	}

	// Populate response object
	response.UpdateTicketStatus.AffectedRow = mutation.UpdatePayment.AffectedRow

	fmt.Printf("Successfully updated payment status for TxRef: %s. Rows affected: %d\n", txref, mutation.UpdatePayment.AffectedRow)
	return response, nil
}

package services
import (
	"cinema-backend/model"
	"context"
	"cinema-backend/utils"
	"github.com/diegosz/go-graphql-client"
	"net/http"
	"os"
	"fmt"
)
// Function to fetch payment data from Hasura using GraphQL
func FetchPaymentData(tx_ref string) (models.TicketFetchOutput, error) {
	// Prepare a query to fetch data
	var query struct {
		PaymentData []models.TicketData `graphql:"ticket(where: {tx_ref: {_eq: $tx_ref}})"`
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
		return models.TicketFetchOutput{}, fmt.Errorf("failed to execute query: %w", err)
	}

	fmt.Println("Query result:", query)
	return models.TicketFetchOutput{
		PaymentData: query.PaymentData,
	}, nil
}
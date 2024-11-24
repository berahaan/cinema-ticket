package utils

import (
	models "cinema-backend/model"
	"context"
	"fmt"
	"github.com/diegosz/go-graphql-client"
	"os"
)

// GetUserByEmail fetches a user by email
func GetUserByEmail(email *string) (response models.GetUserByEmailResponse, err error) {
	client := graphql.NewClient(os.Getenv("GRAPHQL_ENDPOINT"), nil)
	var query struct {
		Users []models.Users `graphql:"users(where: {email: {_eq: $email}})"`
	}
	variables := map[string]interface{}{
		"email": graphql.String(*email),
	}
	err = client.Query(context.Background(), &query, variables)
	if err != nil {
		return response, fmt.Errorf("failed to execute query: %w", err)
	}
	response.Users = query.Users
	if len(response.Users) == 0 {
		return response, nil
	}

	return response, nil
}

package services

import (
	"cinema-backend/model"
	"context"
	"fmt"
	"github.com/diegosz/go-graphql-client"
	"os"
)

func Execute(variables models.SigninArgs) (response models.InsertUserResponse, err error) {
	// first we need to create client
	fmt.Println("Here execute function is called for signup")
	client := graphql.NewClient(os.Getenv("GRAPHQL_ENDPOINT"), nil)
	// Then we need to prepare query to be written which is mutations and variables here
	var mutation struct {
		InsertUsersOne struct {
			ID int `graphql:"id"`
		} `graphql:"insert_users_one(object: {firstname: $firstname, lastname: $lastname, email: $email, password: $password})"`
	}
	variablesMap := map[string]interface{}{
		"firstname": graphql.String(variables.Firstname),
		"lastname":  graphql.String(variables.Lastname),
		"email":     graphql.String(*variables.Email),
		"password":  graphql.String(*variables.Password),
	}
	// now we need to send it to hasura since anyone can create an account we don't need a authorizations here  ....
	err = client.Mutate(context.Background(), &mutation, variablesMap)
	if err != nil {
		return response, fmt.Errorf("failed to execute mutation: %w", err)
	}
	response.Insertusersone.ID = mutation.InsertUsersOne.ID
	return response, nil
}

package services

import (
	"cinema-backend/model"
	"cinema-backend/utils"
	"context"
	"fmt"
	"github.com/diegosz/go-graphql-client"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func InsertAdminUser(email, firstName, lastName, password, role string) (response models.InsertUserResponse, err error) {
	exists, err := utils.GetUserByEmail(&email)
	if err != nil {
		return response, fmt.Errorf("failed to check if user exists: %v", err)
	}
	if len(exists.Users) > 0 {
		fmt.Println("Admin user already exists, skipping insertion.")
		return response, nil
	}
	client := graphql.NewClient(os.Getenv("GRAPHQL_ENDPOINT"), nil)
	var mutation struct {
		InsertAdminUsersOne struct {
			ID int `graphql:"id"`
		} `graphql:"insert_users_one(object: {firstname: $firstname, lastname: $lastname, email: $email, password: $password,role:$Role})"`
	}
	// Hash the password and send it to hasura
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return response, fmt.Errorf("failed to hash password: %v", err)
	}
	variablesMap := map[string]interface{}{
		"firstname": graphql.String(firstName),
		"lastname":  graphql.String(lastName),
		"email":     graphql.String(email),
		"password":  graphql.String(hashedPassword),
		"Role":      graphql.String(role),
	}
	err = client.Mutate(context.Background(), &mutation, variablesMap)
	if err != nil {
		fmt.Println("Error while executing mutations")
	}
	response.Insertusersone.ID = mutation.InsertAdminUsersOne.ID
	return response, nil
}

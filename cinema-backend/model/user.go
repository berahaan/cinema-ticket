package models

import (
	"github.com/diegosz/go-graphql-client"
)

type Users struct {
	ID       int
	Password graphql.String
	Role     graphql.String
}
type EmailCheck struct {
	Email *string `json:"email"`
}
type InsertAdminOutput struct {
	ID int `graphql:"id"`
}

// Define the main response structure.
type GetUserByEmailResponse struct {
	Users []Users
}
type InsertAdminInputPayload struct {
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Password  string  `json:"password"`
	Email     *string `json:"email"`
	Role      string  `json:"role"`
}

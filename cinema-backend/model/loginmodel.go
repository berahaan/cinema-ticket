package models

type LoginArgs struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

type LoginOutput struct {
	AccessToken string `json:"accessToken"`
	ID          int    `json:"id"`
	Role        string `json:"role"`
	Message     string `json:"message"`
	Status      int    `json:"status"`
}
type GraphQLRequest struct {
	Query     string    `json:"query"`
	Variables LoginArgs `json:"variables"`
}

type User struct {
	ID             int    `json:"id"`
	HashedPassword string `json:"password"`
	Role           string `json:"role"`
}

type GraphQLUserData struct {
	Users []User `json:"users"`
}

type GraphQLResponse struct {
	Data   GraphQLUserData `json:"data,omitempty"`
	Errors []GraphQLError  `json:"errors,omitempty"`
}

type GraphQLRequestForEmailCheck struct {
	Query     string              `json:"query"`
	Variables EmailCheckVariables `json:"variables"`
}

type EmailCheckVariables struct {
	Email *string `json:"email"`
}

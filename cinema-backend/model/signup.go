package models

type SigninOutput struct {
	ID int `json:"id"`
}
type InsertUserResponse struct {
	Insertusersone struct {
		ID int `graphql:"id"`
	} `graphql:"insert_users_one"`
}

type SigninArgs struct {
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Email     *string `json:"email"`
	Password  *string `json:"password"`
}

type ActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            SigninArgs             `json:"input"`
}

type SuccessResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

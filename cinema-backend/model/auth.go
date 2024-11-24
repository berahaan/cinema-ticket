package models

type ActionPayloadAuth struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            ValidateTokenArgs      `json:"input"`
}

type ValidateTokenArgs struct {
	Token string `json:"token"`
}

type ValidateOutput struct {
	IsValid bool `json:"isvalid"`
}

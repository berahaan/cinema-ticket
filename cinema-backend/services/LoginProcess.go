package services

import (
	models "cinema-backend/model"
	"cinema-backend/utils"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Login(args models.LoginArgs) (response models.LoginOutput, err error) {
	// Fetch user by email using GetUserByEmail.
	hasuraResponse, err := utils.GetUserByEmail(args.Email)
	if err != nil {
		return models.LoginOutput{}, fmt.Errorf("error while getting user by emails")
	}
	if len(hasuraResponse.Users) == 0 {
		return models.LoginOutput{}, fmt.Errorf("your email account is not registered please create an account and try again")
	}
	user := hasuraResponse.Users[0]
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(*args.Password))
	if err != nil {
		return models.LoginOutput{}, fmt.Errorf("wrong password; please try again")
	}

	userId := user.ID
	fmt.Println("User id here ", userId)
	accessToken, err := utils.GenerateJWT(userId, *args.Email, string(user.Role))
	if err != nil {
		return models.LoginOutput{}, fmt.Errorf("failed to generate access token: %w", err)
	}
	response = models.LoginOutput{
		AccessToken: accessToken,
		Role:        string(user.Role),
		Message:     "Login successful",
		Status:      200,
	}
	return response, nil
}

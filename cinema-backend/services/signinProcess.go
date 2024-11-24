package services

import (
	models "cinema-backend/model"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func SignIn(args models.SigninArgs) (response models.SigninOutput, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*args.Password), bcrypt.DefaultCost)
	if err != nil {
		return response, fmt.Errorf("failed to hash password: %v", err)
	}
	passwordStr := string(hashedPassword)
	args.Password = &passwordStr
	hasuraResponse, err := Execute(args)

	if err != nil {
		fmt.Println("Error during execute ", err)
		return
	}

	response = models.SigninOutput(hasuraResponse.Insertusersone)
	fmt.Println("Response from signin function:", response)
	return response, nil
}

package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(userID int, email string, role string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 8).Unix(), // Access token expiration (24 hours)
		"https://hasura.io/jwt/claims": map[string]interface{}{
			"x-hasura-user-id":       fmt.Sprintf("%d", userID),
			"x-hasura-default-role":  role,
			"x-hasura-allowed-roles": []string{"regular", "admin", "anonymous", "managers"},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

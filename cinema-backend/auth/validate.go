package auth

import (
	models "cinema-backend/model"
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(args models.ValidateTokenArgs) (models.ValidateOutput, error) {
	if args.Token == "" {
		return models.ValidateOutput{IsValid: false}, errors.New("token cannot be empty")
	}

	token, err := jwt.Parse(args.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return models.ValidateOutput{IsValid: false}, nil
	}

	// Check expiration if token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if exp, ok := claims["exp"].(float64); ok {
			// Convert expiration to time and check if it’s expired
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return models.ValidateOutput{IsValid: false}, nil
			}
		}

		return models.ValidateOutput{IsValid: true}, nil
	}

	// If we reach here, something went wrong; return false
	return models.ValidateOutput{IsValid: false}, nil
}

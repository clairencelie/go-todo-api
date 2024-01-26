package helper

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func ValidateJWT(token string) (*jwt.Token, error) {
	errEnvLoad := godotenv.Load("config.env")

	if errEnvLoad != nil {
		return nil, errEnvLoad
	}

	key := os.Getenv("JWT_KEY")

	keyByte := []byte(key)

	validatedToken, errParseToken := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		}

		return keyByte, nil
	})

	if errParseToken != nil {
		return nil, errParseToken
	}

	return validatedToken, nil
}

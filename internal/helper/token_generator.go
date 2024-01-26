package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateJWT(sub string, exp int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iat": time.Now().Unix(),
			"nbf": time.Now().Unix(),
			"iss": "go-todo-restful-api",
			"sub": sub,
			"exp": exp,
		})

	errEnvLoad := godotenv.Load("config.env")

	if errEnvLoad != nil {
		return "", errEnvLoad
	}

	key := os.Getenv("JWT_KEY")

	keyByte := []byte(key)

	tokenStr, errSignToken := token.SignedString(keyByte)

	if errSignToken != nil {
		return "", errSignToken
	}

	return tokenStr, nil
}

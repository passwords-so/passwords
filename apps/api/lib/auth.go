package lib

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const JWTSecret = "secret" // TODO: change this

func GenerateJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // TODO: change this to a longer expiration time or env var
	})

	return token.SignedString([]byte(JWTSecret))
}

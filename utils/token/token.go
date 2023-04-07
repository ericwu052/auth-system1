package token

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(user_id uint) (string, error) {
	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_MINUTES_LIFESPAN"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["auhtorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

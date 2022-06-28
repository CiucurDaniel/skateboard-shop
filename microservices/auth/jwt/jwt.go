package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

type TokenUtil struct {
	secretKey []byte
}

func NewJwtGenerator(secretKey []byte) *TokenUtil {
	return &TokenUtil{secretKey: secretKey}
}

// Generate JWT

func (j *TokenUtil) GetJwtForUserId(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["aud"] = "skateboardshop.ro"
	claims["iss"] = "auth.microservice"
	claims["exp"] = time.Now().Add(time.Hour * 4).Unix()
	claims["uid"] = userId

	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		log.Println("Error generating JWT token")
		return "", err
	}

	return tokenString, nil
}

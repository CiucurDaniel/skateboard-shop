package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

var SecretKey []byte // this will be loaded from env later on

func init() {
	SecretKey = []byte("mysecretkey")
}

// Generate JWT

func GetJwtForUserId(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["aud"] = "skateboardshop.ro"
	claims["iss"] = "auth.microservice"
	claims["exp"] = time.Now().Add(time.Hour * 4).Unix()
	claims["uid"] = userId

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Println("Error generating JWT token")
	}

	return tokenString, nil
}

package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

var SECRET_KEY []byte // this will be loaded from env later on

func init() {
	SECRET_KEY = []byte("mysecretkey")
}

// Generate JWT

func GetJwtForUserId(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["aud"] = "skateboardshop.ro"
	claims["iss"] = "auth.microservice"
	claims["exp"] = time.Now().Add(time.Hour * 4).Unix()
	claims["uid"] = userId

	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Println("Error generating JWT token")
	}

	return tokenString, nil
}

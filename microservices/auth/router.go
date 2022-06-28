package main

import (
	"auth/applogger"
	"auth/data"
	"auth/jwt"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

type AuthController struct {
	jwtGenerator *jwt.TokenUtil
	logger       *applogger.MyLogger
}

var InvalidUserJson = errors.New("Invalid User in Json")

func (c AuthController) requestLogger(router http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		router.ServeHTTP(w, r)

		finishTime := time.Now()
		elapsedTime := finishTime.Sub(startTime)

		// call logger here
		c.logger.LogHttpRequest(w, r, elapsedTime)
	})
}

func (c AuthController) loginHandler(w http.ResponseWriter, r *http.Request) {
	var secretKey []byte

	if os.Getenv("SECRET_KEY") == "" {
		fmt.Println("Not secret key in env")
		secretKey = []byte("mysecretkey") // add it here until we will get it from env
	}

	jwt.NewJwtGenerator(secretKey)

	token, err := c.jwtGenerator.GetJwtForUserId("1")
	if err != nil {
		c.logger.LogError("Failed getting JWT token", err, "loginHandler")
	}
	fmt.Println(fmt.Sprintf("Got token: %v", token))
	fmt.Fprintf(w, token)
}

func (c AuthController) registerHandler(w http.ResponseWriter, r *http.Request) {
	var user data.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, InvalidUserJson.Error(), http.StatusBadRequest)
		return
	}

	err = data.AddUserToDb(user)
	if err != nil {
		http.Error(w, data.UserAlreadyExists.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "User created %+v", user)
}

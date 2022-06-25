package main

import (
	"auth/jwt"
	"github.com/sirupsen/logrus"
)

type AuthService struct {
	Logger       *logrus.Logger
	JwtGenerator *jwt.TokenUtil
}

func NewAuthService(logger *logrus.Logger, jwtGenerator *jwt.TokenUtil) *AuthService {
	return &AuthService{
		Logger:       logger,
		JwtGenerator: jwtGenerator,
	}
}

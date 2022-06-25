package handlers

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type server struct {
	db *gorm.DB
	router *mux.Router
	logger *logrus.Logger
}

// handlers can access the dependencies via the s server variable.
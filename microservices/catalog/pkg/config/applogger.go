package config

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func NewJsonLogger() *logrus.Logger {

	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "log level",
		},
	}
	logger.WithFields(logrus.Fields{"microservice": "Catalog", "author": "Ciucur Daniel"}) // NOT WORKING

	return logger
}
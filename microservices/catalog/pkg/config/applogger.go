package config

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type AppLogger struct {
	logger *logrus.Logger
}

type LogLevel int

const (
	DEBUG LogLevel = iota
	FATAL
	INFO
	ERROR
)

func NewJsonLogger() *AppLogger {

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

	return &AppLogger{logger: logger}

}

func (l AppLogger) LogWithLevel(level LogLevel, message string) {
	commonFields := logrus.Fields{"microservice": "Catalog", "author": "Ciucur Daniel"}

	switch level {
	case DEBUG:
		l.logger.WithFields(commonFields).Debug(message)
	case ERROR:
		l.logger.WithFields(commonFields).Error(message)
	case INFO:
		l.logger.WithFields(commonFields).Info(message)
	case FATAL:
		l.logger.WithFields(commonFields).Fatal(message)
	}
}

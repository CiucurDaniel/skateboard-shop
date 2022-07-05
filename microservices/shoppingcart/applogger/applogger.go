package applogger

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

type MyLogger struct {
	loggerInstance *logrus.Logger
}

func NewAppLogger() *MyLogger {

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

	myLogger := &MyLogger{loggerInstance: logger}
	return myLogger
}

func (l MyLogger) LogHttpRequest(_ http.ResponseWriter, r *http.Request, duration time.Duration) {

	l.loggerInstance.WithFields(
		logrus.Fields{"microservice": "Auth",
			"author":   "Ciucur Daniel",
			"duration": duration,
			"method":   r.Method,
			"uri":      r.RequestURI,
			"host":     r.Host,
		}).Info("Request executed")
}

func (l MyLogger) LogError(message string, err error, methodName string) {
	l.loggerInstance.WithFields(
		logrus.Fields{
			"microservice": "Auth",
			"author":       "Ciucur Daniel",
			"method":       methodName,
			"go error":     err,
		}).Error(message)
}

func (l MyLogger) LogInfo(message string) {
	l.loggerInstance.WithFields(
		logrus.Fields{
			"microservice": "Auth",
			"author":       "Ciucur Daniel",
		}).Info(message)
}

func (l MyLogger) LogFatal(message string) {
	l.loggerInstance.WithFields(
		logrus.Fields{
			"microservice": "Auth",
			"author":       "Ciucur Daniel",
		}).Fatal(message)
}

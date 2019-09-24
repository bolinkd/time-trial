package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// Log returns the logger object
func Log() *logrus.Logger { return log }

// initLogger mainly for debug purposes
func initLogger() {
	logLevel, err := logrus.ParseLevel(os.Getenv("SIO_LL"))
	if err != nil {
		logLevel = logrus.WarnLevel
	}

	log = &logrus.Logger{
		Formatter: new(logrus.TextFormatter),
		Out:       os.Stdout,
		Level:     logLevel,
	}
}

func init() { initLogger() }

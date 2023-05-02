package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// func GetLogger() zerolog.Logger {
// 	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
// 	return logger
// }

var logrusInstance *logrus.Logger

func GetLogger() (logger *logrus.Entry) {
	if logrusInstance != nil {
		return logrusInstance.WithFields(logrus.Fields{})
	}

	dir := "logs"
	fileName := "image-cdn.log"
	path := dir + "/" + fileName

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		_, err = os.Create(path)
		if err != nil {
			panic(err)
		}
	}

	lumberjackLogger := &lumberjack.Logger{
		Filename:  dir + "/" + fileName,
		LocalTime: true,
	}

	logrus.SetOutput(lumberjackLogger)
	logrusInstance = logrus.New()
	logrusInstance.Formatter = &logrus.JSONFormatter{}
	logrusInstance.Level = logrus.DebugLevel
	logrusInstance.Out = lumberjackLogger

	mw := io.MultiWriter(os.Stdout, lumberjackLogger)
	logrusInstance.Out = mw

	return logrusInstance.WithFields(logrus.Fields{})
}

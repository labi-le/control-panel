package utils

import (
	"github.com/sirupsen/logrus"
)

var logger *Logger

func ConfigureLogger(lvl string) {
	levelValid, err := logrus.ParseLevel(lvl)
	if err != nil {
		panic("invalid log level")
	}
	l := logrus.New()
	l.SetLevel(levelValid)
	logger = NewLogger(l)
}

func Log() *Logger {
	return logger
}

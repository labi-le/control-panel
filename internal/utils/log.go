package utils

import (
	"github.com/labi-le/control-panel/pkg"
	"github.com/sirupsen/logrus"
)

var logger *pkg.Logger

func ConfigureLogger(lvl string) {
	levelValid, err := logrus.ParseLevel(lvl)
	if err != nil {
		panic("invalid log level")
	}
	l := logrus.New()
	l.SetLevel(levelValid)
	logger = pkg.NewLogger(l)
}

func Log() *pkg.Logger {
	return logger
}

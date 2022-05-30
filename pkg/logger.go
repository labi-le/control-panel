package pkg

import "github.com/sirupsen/logrus"

type Logger struct {
	*logrus.Logger
}

func NewLogger(logger *logrus.Logger) *Logger {
	return &Logger{Logger: logger}
}

func (l *Logger) Info(s string, i ...interface{}) {
	l.Log(logrus.InfoLevel, s, i)
}

func (l *Logger) Error(s string, i ...interface{}) {
	l.Log(logrus.ErrorLevel, s, i)
}

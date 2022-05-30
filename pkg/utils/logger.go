package utils

import "github.com/sirupsen/logrus"

// Logger
// Compatability with overseer.Logger
// Trick to avoid breaking the interface
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

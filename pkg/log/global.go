package log

import "context"

var GlobalLog = New()

// SetGlobalLog sets global logger
func SetGlobalLog(log Logger) {
	GlobalLog = log
}

// With returns a logger based off the root logger and decorates it with the given context and arguments.
func With(ctx context.Context, args ...interface{}) Logger {
	return GlobalLog.With(ctx, args...)
}

func Debug(args ...interface{}) {
	GlobalLog.Debug(args...)
}

func Info(args ...interface{}) {
	GlobalLog.Info(args...)
}

func Error(args ...interface{}) {
	GlobalLog.Error(args...)
}

func Warn(args ...interface{}) {
	GlobalLog.Warn(args...)
}

func Fatal(args ...interface{}) {
	GlobalLog.Fatal(args...)
}

func Debugf(format string, args ...interface{}) {
	GlobalLog.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	GlobalLog.Infof(format, args...)
}

func Errorf(format string, args ...interface{}) {
	GlobalLog.Errorf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	GlobalLog.Warnf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	GlobalLog.Fatalf(format, args...)
}

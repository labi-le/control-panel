package log

import "context"

type NilLogger struct{}

func (n NilLogger) Warn(_ ...interface{}) {
}

func (n NilLogger) Warnf(_ string, _ ...interface{}) {}

func (n NilLogger) With(_ context.Context, _ ...interface{}) Logger { return n }

func (n NilLogger) Debug(_ ...interface{}) {}

func (n NilLogger) Info(_ ...interface{}) {}

func (n NilLogger) Error(_ ...interface{}) {}

func (n NilLogger) Debugf(_ string, _ ...interface{}) {}

func (n NilLogger) Infof(_ string, _ ...interface{}) {}

func (n NilLogger) Errorf(_ string, _ ...interface{}) {}

func (n NilLogger) Fatal(_ ...interface{}) {}

func (n NilLogger) Fatalf(_ string, _ ...interface{}) {}

package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	slg *zap.Logger
}

func NewLogger(options ...zap.Option) *Logger {
	slg, _ := zap.NewDevelopment(options...)

	return &Logger{
		slg: slg,
	}
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.slg.Sugar().Debugf(template, args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.slg.Sugar().Infof(template, args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.slg.Sugar().Warnf(template, args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.slg.Sugar().Errorf(template, args...)
}

func (l *Logger) Panicf(template string, args ...interface{}) {
	l.slg.Sugar().Panicf(template, args...)
}

func (l *Logger) Sync() error {
	e := l.slg.Sugar().Sync()
	return e
}

package kooperzap

import (
	"fmt"

	"github.com/spotahome/kooper/v2/log"
	"go.uber.org/zap"
)

// Logger is the interface that matches the requirement for kooper logging.
type Logger interface {
	Infof(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	WithKV(log.KV) Logger
}

// KooperLogger makes possible to apply a zap logger to kooper logger.
type KooperLogger struct {
	logger *zap.Logger
}

// NewKooperLogger creates a new object assignable for kooper logger.
func NewKooperLogger(logger *zap.Logger) *KooperLogger {
	return &KooperLogger{logger: logger}
}

// Implementation of different objects and methods to satisfy logging interface.
type KV = map[string]interface{}

func (l *KooperLogger) WithKV(kv log.KV) log.Logger { //nolint:ireturn
	fields := make([]zap.Field, 0, len(kv))
	for k, v := range kv {
		fields = append(fields, zap.Any(k, v))
	}

	return &KooperLogger{logger: l.logger.With(fields...)}
}

func (l *KooperLogger) Infof(format string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, args...))
}

func (l *KooperLogger) Warningf(format string, args ...interface{}) {
	l.logger.Warn(fmt.Sprintf(format, args...))
}

func (l *KooperLogger) Errorf(format string, args ...interface{}) {
	l.logger.Error(fmt.Sprintf(format, args...))
}

func (l *KooperLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debug(fmt.Sprintf(format, args...))
}

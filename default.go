package zapsls

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
	"strings"
)

type defaultLogger struct {
	logger      *zap.Logger
	fieldManger *fieldManger
}

func NewLogger(options ...LoggerOption) Logger {
	defaultTraceID := strings.Replace(uuid.New().String(), "-", "", -1)
	l := &defaultLogger{
		logger: logger,
		fieldManger: &fieldManger{
			traceID: defaultTraceID,
		},
	}
	for _, option := range options {
		option(l)
	}
	return l
}

func (log *defaultLogger) Error(msg string, fields ...zap.Field) {
	fs := log.fieldManger.GetFields(fields...)
	log.logger.Error(msg, fs...)
}

func (log *defaultLogger) Fatal(msg string, fields ...zap.Field) {
	fs := log.fieldManger.GetFields(fields...)
	log.logger.Fatal(msg, fs...)
}

func (log *defaultLogger) Info(msg string, fields ...zap.Field) {
	fs := log.fieldManger.GetFields(fields...)
	log.logger.Info(msg, fs...)
}

func (log *defaultLogger) Warn(msg string, fields ...zap.Field) {
	fs := log.fieldManger.GetFields(fields...)
	log.logger.Warn(msg, fs...)
}

func (log *defaultLogger) Debug(msg string, fields ...zap.Field) {
	fs := log.fieldManger.GetFields(fields...)
	log.logger.Debug(msg, fs...)
}

func (log *defaultLogger) With(field zap.Field) Logger {
	log.fieldManger.AddField(field)
	return log
}

func (log *defaultLogger) GetTraceID() string {
	return log.fieldManger.traceID
}

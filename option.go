package zapsls

type LoggerOption func(*defaultLogger)

func WithTraceID(traceID string) LoggerOption {
	return func(l *defaultLogger) {
		l.fieldManger.traceID = traceID
	}
}

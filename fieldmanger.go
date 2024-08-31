package zapsls

import (
	"go.uber.org/zap"
	"sync"
)

type fieldManger struct {
	traceID string

	m      sync.Mutex
	fields []zap.Field
}

func (fm *fieldManger) AddField(field zap.Field) {
	fm.m.Lock()
	defer fm.m.Unlock()
	fm.fields = append(fm.fields, field)
}

func (fm *fieldManger) GetFields(fields ...zap.Field) []zap.Field {
	uniqMap := map[string]zap.Field{}

	fm.m.Lock()
	for _, f := range fm.fields {
		uniqMap[f.Key] = f
	}
	fm.m.Unlock()

	for _, f := range fields {
		uniqMap[f.Key] = f
	}

	if fm.traceID != "" {
		uniqMap[TraceIDKey] = zap.String(TraceIDKey, fm.traceID)
	}

	ret := []zap.Field{}
	for _, f := range uniqMap {
		ret = append(ret, f)
	}
	return ret
}

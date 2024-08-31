package zapsls

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

var logger *zap.Logger

func InitLogger(config *Config) {
	err := config.Validate()
	if err != nil {
		panic(err)
	}

	if logger != nil {
		return
	}

	w := newDefaultWriter(config)
	ws := zapcore.AddSync(w)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, ws, config.LogLevel)
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	log.Println("初始化 SLS 完毕")
}

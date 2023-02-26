package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// AccessLog and Log are package level variables, every program should access logging function through these variables
var (
	AccessLog *zap.Logger
	Log       *zap.Logger
)

func SetupAccessLogger(outputPaths []string) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = outputPaths
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	AccessLog = logger
}

func SetupDebugLogger(outputPaths []string) {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = outputPaths
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	Log = logger
}

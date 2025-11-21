package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var myLogger *zap.Logger
var atom = zap.NewAtomicLevelAt(zap.DebugLevel)

func init() {
	config := zap.Config{
		Level:       atom,
		Development: false,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "msg",
			LevelKey:      "level",
			TimeKey:       "time",
			CallerKey:     "caller",
			StacktraceKey: "stacktrace",
			EncodeTime:    zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	var err error
	myLogger, err = config.Build()
	if err != nil {
		panic(err)
	}
	myLogger = myLogger.WithOptions(zap.AddCallerSkip(1))
}

func SetLevel(level string) {
	parseLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	atom.SetLevel(parseLevel)
}

func Debug(msg string, fields ...zap.Field) {
	myLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	myLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	myLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	myLogger.Error(msg, fields...)
}

package main

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func formatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%d%02d%02d_%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
}
func FormateLog(args []interface{}) *zap.Logger {
	log := logger.With(ToJsonData(args))
	return log
}
func Debug(msg string, args ...interface{}) {
	FormateLog(args).Sugar().Debugf(msg)
}
func ToJsonData(args []interface{}) zap.Field {
	det := make([]string, 0)
	if len(args) > 0 {
		for _, v := range args {
			det = append(det, fmt.Sprintf("%+v", v))
		}
	}
	zap := zap.Any("detail", det)
	return zap
}
func InitZapLog() {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "t",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "trace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     formatEncodeTime,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"zap.log"},
		ErrorOutputPaths: []string{"zap.log"},
		InitialFields: map[string]interface{}{
			"app": "test",
		},
	}
	var err error
	logger, err = cfg.Build()
	if err != nil {
		panic("log init fail:" + err.Error())
	}
}
func main() {
	InitZapLog()
	defer logger.Sync()
	a := []string{"test", "hello", "world"}
	Debug("output", a)
}

package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var logger *zap.Logger

// 初始化配置
func InitZapLog() {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			MessageKey:     "M",
			StacktraceKey:  "S",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"./zap.log"},
		ErrorOutputPaths: []string{"./zap_error.log"},
		InitialFields: map[string]interface{}{
			"app": "myApp",
		},
	}
	var err error
	logger, err = cfg.Build()
	if err != nil {
		panic("log init fail:" + err.Error())
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%d%02d%02d_%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
}

func Info(msg string, args ...interface{}) {
	FormateLog(args).Sugar().Info(msg)
}

func Debug(msg string, args ...interface{}) {
	FormateLog(args).Sugar().Debugf(msg)
}

func Error(msg string, args ...interface{}) {
	FormateLog(args).Sugar().Error(msg)
}

func FormateLog(args []interface{}) *zap.Logger {
	log := logger.With(ToJsonData(args))
	return log
}

func ToJsonData(args []interface{}) zap.Field {
	det := make([]string, 0)
	if len(args) > 0 {
		for _, v := range args {
			det = append(det, fmt.Sprintf("%+v", v))
		}
	}
	return zap.Any("detail", det)
}

func main() {
	InitZapLog()
	// 程序退出前调用，刷新所有缓冲的日志
	defer logger.Sync()

	a := []string{"hello", "world"}
	Error("output", a)
}

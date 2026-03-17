package config

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.SugaredLogger

func InitLogger() error {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	level := getLogLevel()

	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	Logger = logger.Sugar()
	return nil
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	var filename string
	var maxSize, maxBackups, maxAge int
	var compress bool

	if GlobalConfig != nil {
		filename = GlobalConfig.Log.Filename
		maxSize = GlobalConfig.Log.MaxSize
		maxBackups = GlobalConfig.Log.MaxBackups
		maxAge = GlobalConfig.Log.MaxAge
		compress = GlobalConfig.Log.Compress
	}

	if filename == "" {
		filename = "logs/blog.log"
	}
	if maxSize == 0 {
		maxSize = 10
	}
	if maxBackups == 0 {
		maxBackups = 5
	}
	if maxAge == 0 {
		maxAge = 30
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
}

func getLogLevel() zapcore.Level {
	if GlobalConfig == nil || GlobalConfig.Log.Level == "" {
		return zapcore.InfoLevel
	}

	switch GlobalConfig.Log.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func Info(args ...interface{}) {
	Logger.Info(args...)
}

func Error(args ...interface{}) {
	Logger.Error(args...)
}

func Warn(args ...interface{}) {
	Logger.Warn(args...)
}

func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

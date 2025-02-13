package logger

import (
	"log"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	once   sync.Once
)

func LogInit(level string) {
	once.Do(func() {
		defer func() {
			_ = logger.Sync()
		}()

		var zapLevel zapcore.Level
		switch level {
		case "debug":
			zapLevel = zapcore.DebugLevel
		case "info":
			zapLevel = zapcore.InfoLevel
		case "warn":
			zapLevel = zapcore.WarnLevel
		case "error":
			zapLevel = zapcore.ErrorLevel
		case "fatal":
			zapLevel = zapcore.FatalLevel
		default:
			zapLevel = zapcore.InfoLevel
		}

		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeName:     zapcore.FullNameEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
		}

		fileWriteSyncer := zapcore.Lock(zapcore.AddSync(mustOpenFile("./name.log")))
		consoleWriteSyncer := zapcore.AddSync(os.Stdout)

		core := zapcore.NewTee(
			zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), fileWriteSyncer, zapLevel),
			zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), consoleWriteSyncer, zapLevel),
		)

		logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	})
}

func mustOpenFile(path string) zapcore.WriteSyncer {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Failed to open log file: %v\n", err)
		return zapcore.AddSync(os.Stdout)
	}
	return zapcore.AddSync(file)
}

func ensureLoggerInitialized() {
	if logger == nil {
		log.Println("Logger not initialized. Initializing with default settings.")
		LogInit("info")
	}
}

func GetLogger() *zap.Logger {
	ensureLoggerInitialized()
	return logger
}

func Info(message string, fields ...zap.Field) {
	GetLogger().Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	GetLogger().Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	GetLogger().Error(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	GetLogger().Warn(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	GetLogger().Fatal(message, fields...)
}

func DPanic(message string, fields ...zap.Field) {
	GetLogger().DPanic(message, fields...)
}

func Panic(message string, fields ...zap.Field) {
	GetLogger().Panic(message, fields...)
}

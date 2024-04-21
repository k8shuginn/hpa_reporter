package logger

import (
	"go.uber.org/zap"
)

var (
	// writer is the global logger
	writer *zap.Logger
)

// Debug logs a message at DebugLevel. The message includes any fields passed
func Debug(msg string, fields ...zap.Field) {
	writer.Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
func Info(msg string, fields ...zap.Field) {
	writer.Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
func Warn(msg string, fields ...zap.Field) {
	writer.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
func Error(msg string, fields ...zap.Field) {
	writer.Error(msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
func Fatal(msg string, fields ...zap.Field) {
	writer.Fatal(msg, fields...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
func Panic(msg string, fields ...zap.Field) {
	writer.Panic(msg, fields...)
}

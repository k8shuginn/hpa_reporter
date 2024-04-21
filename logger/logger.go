package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

const (
	DefaultDir string = "log"
)

func TestSetLogger() {
	dir, _ := os.Getwd()

	SetLogger("test-logger",
		WithLogPath(filepath.Join(dir, DefaultDir)),
		WithLogLocalTime(true),
		WithLogLevel("debug"),
	)
}

// SetLogger initializes the global logger
func SetLogger(name string, opts ...Option) {
	cfg := fromOptions(name, opts...)
	writer = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.NewMultiWriteSyncer(append([]zapcore.WriteSyncer{zapcore.AddSync(os.Stdout)}, zapcore.AddSync(&cfg.logger))...),
		cfg.level,
	))
}

package logger

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"path/filepath"
	"strings"
)

const (
	DefaultPath  string = "/var/log/k8shuginn/"
	LogExtension string = ".log"
)

type config struct {
	name   string
	logger lumberjack.Logger
	level  zapcore.Level
}

// defaultConfig returns a default configuration for a logger
func defaultConfig(name string) config {
	return config{
		name: name,
		logger: lumberjack.Logger{
			Filename:   filepath.Join(DefaultPath, name+LogExtension),
			MaxSize:    100,   // MB
			MaxAge:     7,     // days
			MaxBackups: 10,    // file count
			LocalTime:  false, // UTC
			Compress:   true,
		},
		level: zapcore.InfoLevel,
	}
}

// fromOptions returns a configuration for a logger based on the provided options
func fromOptions(name string, opts ...Option) *config {
	cfg := defaultConfig(name)
	for _, opt := range opts {
		opt(&cfg)
	}

	return &cfg
}

// Option is a function that modifies a configuration
type Option func(*config)

// WithLogPath sets the path for the log file
func WithLogPath(path string) Option {
	return func(c *config) {
		if path != "" {
			c.logger.Filename = filepath.Join(path, c.name+LogExtension)
		}
	}
}

// WithLogMaxSize sets the maximum size of the log file
func WithLogMaxSize(size int) Option {
	return func(c *config) {
		if size > 1 {
			c.logger.MaxSize = size
		}
	}
}

// WithLogMaxAge sets the maximum age of the log file
func WithLogMaxAge(age int) Option {
	return func(c *config) {
		if age > 1 {
			c.logger.MaxAge = age
		}
	}
}

// WithLogMaxBackups sets the maximum number of backups for the log file
func WithLogMaxBackups(count int) Option {
	return func(c *config) {
		if count > 1 {
			c.logger.MaxBackups = count
		}
	}
}

// WithLogLocalTime sets the local time for the log file
func WithLogLocalTime(localTime bool) Option {
	return func(c *config) {
		c.logger.LocalTime = localTime
	}
}

// WithLogCompress sets the compression for the log file
func WithLogCompress(compress bool) Option {
	return func(c *config) {
		c.logger.Compress = compress
	}
}

// WithLogLevel sets the log level for the logger
func WithLogLevel(level string) Option {
	var lv zapcore.Level

	switch strings.ToUpper(level) {
	case "DEBUG", "DEB":
		lv = zapcore.DebugLevel
	case "INFO", "INF":
		lv = zapcore.InfoLevel
	case "WARN", "WARNING":
		lv = zapcore.WarnLevel
	case "ERROR", "ERR":
		lv = zapcore.ErrorLevel
	case "DPANIC", "PANIC", "PNC":
		lv = zapcore.PanicLevel
	case "FATAL", "FTL":
		lv = zapcore.FatalLevel
	default:
		lv = zapcore.InfoLevel
	}

	return func(c *config) {
		c.level = lv
	}
}

package main

import (
	"github.com/k8shuginn/hpa_reporter/cmd/hpa_reporter/app"
	"github.com/k8shuginn/hpa_reporter/logger"
	"go.uber.org/zap"
	"os"
	"strconv"
)

const (
	EnvLogPath    = "LOG_PATH"
	EnvLogMaxSize = "LOG_MAX_SIZE"
	EnvLogAge     = "LOG_AGE"
	EnvLogBackups = "LOG_BACKUPS"
	EnvLogLevel   = "LOG_LEVEL"
)

func init() {
	path := os.Getenv(EnvLogPath)
	maxSize, _ := strconv.Atoi(os.Getenv(EnvLogMaxSize))
	age, _ := strconv.Atoi(os.Getenv(EnvLogAge))
	backups, _ := strconv.Atoi(os.Getenv(EnvLogBackups))
	level := os.Getenv(EnvLogLevel)

	logger.SetLogger(
		app.Name,
		logger.WithLogPath(path),
		logger.WithLogMaxSize(maxSize),
		logger.WithLogMaxAge(age),
		logger.WithLogMaxBackups(backups),
		logger.WithLogLevel(level),
	)

	logger.Info("logger initialized", zap.String("app", app.Name), zap.String("path", path), zap.String("level", level))
}

func main() {
	reporter := app.NewApp()
	if err := reporter.Init(); err != nil {
		logger.Fatal("failed to initialize app", zap.Error(err))
	}

	reporter.Execute()
}

package app

import (
	"fmt"
	"github.com/alecthomas/kingpin/v2"
	"github.com/k8shuginn/hpa_reporter/cmd/hpa_reporter/app/collector"
	"github.com/k8shuginn/hpa_reporter/cmd/hpa_reporter/app/config"
	"github.com/k8shuginn/hpa_reporter/cmd/hpa_reporter/app/reporter"
	"github.com/k8shuginn/hpa_reporter/logger"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

var (
	ConfigPath = kingpin.Flag("app.config", "config file path.").Default(config.DefaultConfigPath).String()
	Name       = "hpa_reporter"
)

func init() {
	kingpin.Parse()
}

type App struct {
	appConfig *config.AppConfig
	rh        *reporter.Handler
	ch        *collector.Handler
}

func NewApp() *App {
	return &App{}
}

func (a *App) Init() error {
	var err error

	// load config
	a.appConfig, err = config.LoadConfig(*ConfigPath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	logger.Info("config loaded", zap.Any("config", a.appConfig))

	// create reporter handler
	a.rh, err = reporter.NewReporterHandler(a.appConfig.Reporters)
	if err != nil {
		return fmt.Errorf("failed to create reporter handler: %w", err)
	}

	// create collector handler
	a.ch, err = collector.NewCollectorHandler(a.rh, a.appConfig.Hpa)
	if err != nil {
		return fmt.Errorf("failed to create collector handler: %w", err)
	}

	return nil
}

func (a *App) Execute() {
	logger.Info(Name + " is started ... ")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	a.ch.Run()

	select {
	case sig := <-sigChan:
		a.Shutdown()
		logger.Info(Name+" is shutdown", zap.String("signal", sig.String()))
	}
}

func (a *App) Shutdown() {
	a.ch.Shutdown()
	logger.Info(Name + " is stopped ... ")

	a.rh.Shutdown()
	logger.Info("reporter is stopped ... ")
}

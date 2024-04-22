package reporter

import (
	"fmt"
	"github.com/k8shuginn/hpa_reporter/cmd/hpa-reporter/app/config"
	"github.com/k8shuginn/hpa_reporter/cmd/hpa-reporter/app/message"
	"github.com/k8shuginn/hpa_reporter/cmd/hpa-reporter/app/reporter/slack"
	"github.com/k8shuginn/hpa_reporter/cmd/hpa-reporter/app/reporter/stdout"
)

type Reporter interface {
	Report(msg *message.Data)
}

// Handler is reporter handler
type Handler struct {
	shutdown     chan struct{}
	ReporterList []Reporter
}

// NewReporterHandler creates a new reporter handler
func NewReporterHandler(reporterConfig config.ReporterConfig) (*Handler, error) {
	h := &Handler{
		shutdown: make(chan struct{}),
	}

	for _, cfg := range reporterConfig.Stdout {
		rep, _ := stdout.CreateReporter(cfg, h.shutdown)
		h.ReporterList = append(h.ReporterList, rep)
	}

	for _, cfg := range reporterConfig.Slack {
		rep, err := slack.CreateReporter(cfg, h.shutdown)
		if err != nil {
			h.Shutdown()
			return nil, fmt.Errorf("failed to create slack reporter: %w", err)
		}
		h.ReporterList = append(h.ReporterList, rep)
	}

	return h, nil
}

// Report sends message to all reporters
func (h *Handler) Report(msg *message.Data) {
	for _, rep := range h.ReporterList {
		go rep.Report(msg)
	}
}

// Shutdown stops all reporters
func (h *Handler) Shutdown() {
	close(h.shutdown)
}

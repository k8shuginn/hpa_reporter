package collector

import (
	"fmt"
	"github.com/k8shuginn/hpa_reporter/cmd/hpa-reporter/app/config"
	"github.com/k8shuginn/hpa_reporter/cmd/hpa-reporter/app/message"
	"github.com/k8shuginn/hpa_reporter/k8s"
	"github.com/k8shuginn/hpa_reporter/logger"
	"go.uber.org/zap"
	"os"
)

const (
	EnvKubeConfig = "KUBECONFIG"
)

type Reporter interface {
	Report(*message.Data)
}

// Handler is kubernetes hpa event handler.
type Handler struct {
	reporter  Reporter
	client    *k8s.Client
	hpaTarget map[string]int32

	OnAddFunc    // unused
	OnUpdateFunc // used
	OnDeleteFunc // unused
}

// NewCollectorHandler is a constructor that creates a new handler.
func NewCollectorHandler(reporter Reporter, configs []config.HpaConfig) (*Handler, error) {
	h := &Handler{
		reporter:  reporter,
		hpaTarget: make(map[string]int32),
	}

	// set hpa target
	for _, cfg := range configs {
		k := cfg.Namespace + "/" + cfg.Name
		h.hpaTarget[k] = cfg.Threshold
	}

	// create k8s client
	client, err := k8s.NewClient(h, os.Getenv(EnvKubeConfig))
	if err != nil {
		return nil, fmt.Errorf("[collector] failed to create k8s client: %w", err)
	}
	h.client = client

	// set hpa version
	version := h.client.GetHPAVersion()
	switch version {
	case "v1":
		h.OnAddFunc = h.v1Add
		h.OnUpdateFunc = h.v1Update
		h.OnDeleteFunc = h.v1Delete
	case "v2":
		h.OnAddFunc = h.v2Add
		h.OnUpdateFunc = h.v2Update
		h.OnDeleteFunc = h.v2Delete
	case "v2beta1":
		h.OnAddFunc = h.v2beta1Add
		h.OnUpdateFunc = h.v2beta1Update
		h.OnDeleteFunc = h.v2beta1Delete
	case "v2beta2":
		h.OnAddFunc = h.v2beta2Add
		h.OnUpdateFunc = h.v2beta2Update
		h.OnDeleteFunc = h.v2beta2Delete
	default:
		return nil, fmt.Errorf("[collector] unsupported hpa version: %s", h.client.GetHPAVersion())
	}
	logger.Info("[collector] k8s client created", zap.String("hpa version", version))

	return h, nil
}

// Run is a method that starts the handler.
func (h *Handler) Run() {
	h.client.Start()
	logger.Info("[collector] is started ... ")
}

// Shutdown is a method that stops the handler.
func (h *Handler) Shutdown() {
	h.client.Stop()
	logger.Info("[collector] is stopped ... ")
}

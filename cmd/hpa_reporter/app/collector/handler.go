package collector

import (
	"fmt"
	"github.com/k8shuginn/hpa_reporter/cmd/hpa_reporter/app/config"
	"github.com/k8shuginn/hpa_reporter/cmd/hpa_reporter/app/message"
	"github.com/k8shuginn/hpa_reporter/k8s"
	"github.com/k8shuginn/hpa_reporter/logger"
	"go.uber.org/zap"
	v1 "k8s.io/api/autoscaling/v1"
	v2 "k8s.io/api/autoscaling/v2"
	"k8s.io/api/autoscaling/v2beta1"
	"k8s.io/api/autoscaling/v2beta2"
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
		h.OnUpdateFunc = h.v1Update
	case "v2":
		h.OnUpdateFunc = h.v2Update
	case "v2beta1":
		h.OnUpdateFunc = h.v2beta1Update
	case "v2beta2":
		h.OnUpdateFunc = h.v2beta2Update
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

// OnAdd is a method that handles the add event.
func (h *Handler) v1Update(_, obj interface{}) {
	object, ok := obj.(*v1.HorizontalPodAutoscaler)
	if !ok {
		logger.Error("[HpaEvent] v1Update type assertion error")
		return
	}

	k := object.Namespace + "/" + object.Name
	threshold, ok := h.hpaTarget[k]
	if !ok {
		return
	}

	if object.Status.CurrentReplicas >= object.Spec.MaxReplicas {
		h.reporter.Report(&message.Data{
			Level:           message.LevelCritical,
			Name:            object.Name,
			Namespace:       object.Namespace,
			CurrentReplicas: object.Status.CurrentReplicas,
			MaxReplicas:     object.Spec.MaxReplicas,
		})
	} else if object.Status.CurrentReplicas >= threshold {
		h.reporter.Report(&message.Data{
			Level:           message.LevelWarning,
			Name:            object.Name,
			Namespace:       object.Namespace,
			CurrentReplicas: object.Status.CurrentReplicas,
			MaxReplicas:     object.Spec.MaxReplicas,
		})
	}
}

// v2Update is a method that handles the v2.HorizontalPodAutoscaler update event.
func (h *Handler) v2Update(_, obj interface{}) {
	object, ok := obj.(*v2.HorizontalPodAutoscaler)
	if !ok {
		logger.Error("[HpaEvent] v2Update type assertion error")
		return
	}

	k := object.Namespace + "/" + object.Name
	threshold, ok := h.hpaTarget[k]
	if !ok {
		return
	}

	if object.Status.CurrentReplicas >= object.Spec.MaxReplicas {
		h.reporter.Report(&message.Data{
			Level:           message.LevelCritical,
			Name:            object.Name,
			Namespace:       object.Namespace,
			CurrentReplicas: object.Status.CurrentReplicas,
			MaxReplicas:     object.Spec.MaxReplicas,
		})
	} else if object.Status.CurrentReplicas >= threshold {
		h.reporter.Report(&message.Data{
			Level:           message.LevelWarning,
			Name:            object.Name,
			Namespace:       object.Namespace,
			CurrentReplicas: object.Status.CurrentReplicas,
			MaxReplicas:     object.Spec.MaxReplicas,
		})
	}
}

// v2beta1Update is a method that handles the v2beta1.HorizontalPodAutoscaler update event.
func (h *Handler) v2beta1Update(_, obj interface{}) {
	object, ok := obj.(*v2beta1.HorizontalPodAutoscaler)
	if !ok {
		logger.Error("[HpaEvent] v2beta1Update type assertion error")
		return
	}

	k := object.Namespace + "/" + object.Name
	threshold, ok := h.hpaTarget[k]
	if !ok {
		return
	}

	if object.Status.CurrentReplicas >= object.Spec.MaxReplicas {
		h.reporter.Report(&message.Data{
			Level:           message.LevelCritical,
			Name:            object.Name,
			Namespace:       object.Namespace,
			CurrentReplicas: object.Status.CurrentReplicas,
			MaxReplicas:     object.Spec.MaxReplicas,
		})
	} else if object.Status.CurrentReplicas >= threshold {
		h.reporter.Report(&message.Data{
			Level:           message.LevelWarning,
			Name:            object.Name,
			Namespace:       object.Namespace,
			CurrentReplicas: object.Status.CurrentReplicas,
			MaxReplicas:     object.Spec.MaxReplicas,
		})
	}
}

// v2beta2Update is a method that handles v2beta2.HorizontalPodAutoscaler events
func (h *Handler) v2beta2Update(_, obj interface{}) {
	object, ok := obj.(*v2beta2.HorizontalPodAutoscaler)
	if !ok {
		logger.Error("[HpaEvent] v2beta2Update type assertion error")
		return
	}

	k := object.Namespace + "/" + object.Name
	threshold, ok := h.hpaTarget[k]
	if !ok {
		return
	}

	if object.Status.CurrentReplicas >= object.Spec.MaxReplicas {
		h.reporter.Report(&message.Data{
			Level:           message.LevelCritical,
			Name:            object.Name,
			Namespace:       object.Namespace,
			CurrentReplicas: object.Status.CurrentReplicas,
			MaxReplicas:     object.Spec.MaxReplicas,
		})

	} else if object.Status.CurrentReplicas >= threshold {
		h.reporter.Report(&message.Data{
			Level:           message.LevelWarning,
			Name:            object.Name,
			Namespace:       object.Namespace,
			CurrentReplicas: object.Status.CurrentReplicas,
			MaxReplicas:     object.Spec.MaxReplicas,
		})
	}
}

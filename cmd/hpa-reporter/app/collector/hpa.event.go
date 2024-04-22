package collector

import (
	"github.com/k8shuginn/hpa_reporter/cmd/hpa-reporter/app/message"
	"github.com/k8shuginn/hpa_reporter/logger"
	"go.uber.org/zap"
	v1 "k8s.io/api/autoscaling/v1"
	v2 "k8s.io/api/autoscaling/v2"
	v2beta1 "k8s.io/api/autoscaling/v2beta1"
	v2beta2 "k8s.io/api/autoscaling/v2beta2"
)

// Handler is kubernetes hpa event handler.
func (h *Handler) v1Add(obj interface{}, _ bool) {
	object, ok := obj.(*v1.HorizontalPodAutoscaler)
	if !ok {
		logger.Error("[HpaEvent] v1Add type assertion error")
		return
	}

	k := object.Namespace + "/" + object.Name
	_, ok = h.hpaTarget[k]
	if !ok {
		return
	}

	logger.Debug("[HpaEvent] v1Add hpa detected", zap.String("name", object.Name), zap.String("namespace", object.Namespace))
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
			Time:            object.Status.LastScaleTime.Format("2006-01-02 15:04:05"),
			Level:           message.LevelCritical,
			Name:            object.Name,
			Namespace:       object.Namespace,
			CurrentReplicas: object.Status.CurrentReplicas,
			MaxReplicas:     object.Spec.MaxReplicas,
		})
	} else if object.Status.CurrentReplicas >= threshold {
		h.reporter.Report(&message.Data{
			Time:            object.Status.LastScaleTime.Format("2006-01-02 15:04:05"),
			Level:           message.LevelWarning,
			Name:            object.Name,
			Namespace:       object.Namespace,
			CurrentReplicas: object.Status.CurrentReplicas,
			MaxReplicas:     object.Spec.MaxReplicas,
		})
	}
}

// v1Delete is a method that handles the v1.HorizontalPodAutoscaler delete event.
func (h *Handler) v1Delete(obj interface{}) {
	object, ok := obj.(*v1.HorizontalPodAutoscaler)
	if !ok {
		logger.Error("[HpaEvent] v1Delete type assertion error")
		return
	}

	k := object.Namespace + "/" + object.Name
	_, ok = h.hpaTarget[k]
	if !ok {
		return
	}

	logger.Debug("[HpaEvent] v1Delete hpa disappeared", zap.String("name", object.Name), zap.String("namespace", object.Namespace))
}

// v2Add is a method that handles the v2.HorizontalPodAutoscaler add event.
func (h *Handler) v2Add(obj interface{}, _ bool) {
	object, ok := obj.(*v2.HorizontalPodAutoscaler)
	if !ok {
		logger.Error("[HpaEvent] v2Add type assertion error")
		return
	}

	k := object.Namespace + "/" + object.Name
	_, ok = h.hpaTarget[k]
	if !ok {
		return
	}

	logger.Debug("[HpaEvent] v2Add hpa detected", zap.String("name", object.Name), zap.String("namespace", object.Namespace))
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
			Time:            object.Status.LastScaleTime.Format("2006-01-02 15:04:05"),
			Level:           message.LevelCritical,
			Name:            object.Name,
			Namespace:       object.Namespace,
			CurrentReplicas: object.Status.CurrentReplicas,
			MaxReplicas:     object.Spec.MaxReplicas,
		})
	} else if object.Status.CurrentReplicas >= threshold {
		h.reporter.Report(&message.Data{
			Time:            object.Status.LastScaleTime.Format("2006-01-02 15:04:05"),
			Level:           message.LevelWarning,
			Name:            object.Name,
			Namespace:       object.Namespace,
			CurrentReplicas: object.Status.CurrentReplicas,
			MaxReplicas:     object.Spec.MaxReplicas,
		})
	}
}

// v2Delete is a method that handles the v2.HorizontalPodAutoscaler delete event.
func (h *Handler) v2Delete(obj interface{}) {
	object, ok := obj.(*v2.HorizontalPodAutoscaler)
	if !ok {
		logger.Error("[HpaEvent] v2Delete type assertion error")
		return
	}

	k := object.Namespace + "/" + object.Name
	_, ok = h.hpaTarget[k]
	if !ok {
		return
	}

	logger.Debug("[HpaEvent] v2Delete hpa disappeared", zap.String("name", object.Name), zap.String("namespace", object.Namespace))
}

// v2beta1Add is a method that handles the v2beta1.HorizontalPodAutoscaler add event.
func (h *Handler) v2beta1Add(obj interface{}, _ bool) {
	object, ok := obj.(*v2beta1.HorizontalPodAutoscaler)
	if !ok {
		logger.Error("[HpaEvent] v2beta1Add type assertion error")
		return
	}

	k := object.Namespace + "/" + object.Name
	_, ok = h.hpaTarget[k]
	if !ok {
		return
	}

	logger.Debug("[HpaEvent] v2beta1Add hpa detected", zap.String("name", object.Name), zap.String("namespace", object.Namespace))
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
			Time:            object.Status.LastScaleTime.Format("2006-01-02 15:04:05"),
			Level:           message.LevelCritical,
			Name:            object.Name,
			Namespace:       object.Namespace,
			CurrentReplicas: object.Status.CurrentReplicas,
			MaxReplicas:     object.Spec.MaxReplicas,
		})
	} else if object.Status.CurrentReplicas >= threshold {
		h.reporter.Report(&message.Data{
			Time:            object.Status.LastScaleTime.Format("2006-01-02 15:04:05"),
			Level:           message.LevelWarning,
			Name:            object.Name,
			Namespace:       object.Namespace,
			CurrentReplicas: object.Status.CurrentReplicas,
			MaxReplicas:     object.Spec.MaxReplicas,
		})
	}
}

// v2beta1Delete is a method that handles the v2beta1.HorizontalPodAutoscaler delete event.
func (h *Handler) v2beta1Delete(obj interface{}) {
	object, ok := obj.(*v2beta1.HorizontalPodAutoscaler)
	if !ok {
		logger.Error("[HpaEvent] v2beta1Delete type assertion error")
		return
	}

	k := object.Namespace + "/" + object.Name
	_, ok = h.hpaTarget[k]
	if !ok {
		return
	}

	logger.Debug("[HpaEvent] v2beta1Delete hpa disappeared", zap.String("name", object.Name), zap.String("namespace", object.Namespace))

}

// v2beta2Add is a method that handles v2beta2.HorizontalPodAutoscaler events
func (h *Handler) v2beta2Add(obj interface{}, _ bool) {
	object, ok := obj.(*v2beta2.HorizontalPodAutoscaler)
	if !ok {
		logger.Error("[HpaEvent] v2beta2Add type assertion error")
		return
	}

	k := object.Namespace + "/" + object.Name
	_, ok = h.hpaTarget[k]
	if !ok {
		return
	}

	logger.Debug("[HpaEvent] v2beta2Add hpa detected", zap.String("name", object.Name), zap.String("namespace", object.Namespace))

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
			Time:            object.Status.LastScaleTime.Format("2006-01-02 15:04:05"),
			Level:           message.LevelCritical,
			Name:            object.Name,
			Namespace:       object.Namespace,
			CurrentReplicas: object.Status.CurrentReplicas,
			MaxReplicas:     object.Spec.MaxReplicas,
		})

	} else if object.Status.CurrentReplicas >= threshold {
		h.reporter.Report(&message.Data{
			Time:            object.Status.LastScaleTime.Format("2006-01-02 15:04:05"),
			Level:           message.LevelWarning,
			Name:            object.Name,
			Namespace:       object.Namespace,
			CurrentReplicas: object.Status.CurrentReplicas,
			MaxReplicas:     object.Spec.MaxReplicas,
		})
	}
}

// v2beta2Delete is a method that handles v2beta2.HorizontalPodAutoscaler events
func (h *Handler) v2beta2Delete(obj interface{}) {
	object, ok := obj.(*v2beta2.HorizontalPodAutoscaler)
	if !ok {
		logger.Error("[HpaEvent] v2beta2Delete type assertion error")
		return
	}

	k := object.Namespace + "/" + object.Name
	_, ok = h.hpaTarget[k]
	if !ok {
		return
	}

	logger.Debug("[HpaEvent] v2beta2Delete hpa disappeared", zap.String("name", object.Name), zap.String("namespace", object.Namespace))
}

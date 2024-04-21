package k8s

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"strings"
)

const (
	ResourceGroup = "autoscaling"
	ResourceName  = "horizontalpodautoscalers"
)

func (c *Client) initClientSet(kubeConfig string) error {
	var clientCfg *rest.Config
	var err error

	if kubeConfig != "" {
		// out of cluster
		clientCfg, err = clientcmd.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			return fmt.Errorf("failed to create client config: %w", err)
		}
	} else {
		// in cluster
		clientCfg, err = rest.InClusterConfig()
		if err != nil {
			return fmt.Errorf("failed to create client config: %w", err)
		}
	}

	c.cs, err = kubernetes.NewForConfig(clientCfg)
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}

	return nil
}

func (c *Client) hpaShardIndexInformer(version string) error {
	switch version {
	case "v1":
		c.hpaSii = c.iFactory.Autoscaling().V1().HorizontalPodAutoscalers().Informer()
	case "v2":
		c.hpaSii = c.iFactory.Autoscaling().V2().HorizontalPodAutoscalers().Informer()
	case "v2beta1":
		c.hpaSii = c.iFactory.Autoscaling().V2beta1().HorizontalPodAutoscalers().Informer()
	case "v2beta2":
		c.hpaSii = c.iFactory.Autoscaling().V2beta2().HorizontalPodAutoscalers().Informer()
	default:
		return fmt.Errorf("hpa unsupported version : %s", version)
	}

	return nil
}

func (c *Client) initShardIndexInformer(hpaEventHandler cache.ResourceEventHandler) error {
	apiResources, err := c.cs.Discovery().ServerPreferredResources()
	if err != nil {
		return fmt.Errorf("failed to get server preferred resources: %v", err)
	}

LOOP:
	for _, list := range apiResources {
		groupVersion := strings.Split(list.GroupVersion, "/")
		if groupVersion[0] == ResourceGroup {
			for _, apiResource := range list.APIResources {
				if apiResource.Name == ResourceName {
					c.hpaVersion = groupVersion[1]
					if err = c.hpaShardIndexInformer(groupVersion[1]); err != nil {
						return fmt.Errorf("failed to create shard index informer: %v", err)
					}

					break LOOP
				}
			}
		}
	}

	if c.hpaVersion == "" || c.hpaSii == nil {
		return fmt.Errorf("failed to create shard index informer: hpaVersion is empty")
	}

	c.hpaReg, err = c.hpaSii.AddEventHandler(hpaEventHandler)
	if err != nil {
		return fmt.Errorf("failed to add event handler: %v", err)
	}

	return nil
}

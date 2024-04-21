package k8s

import (
	"fmt"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

type Client struct {
	shutdown chan struct{}
	cs       *kubernetes.Clientset
	iFactory informers.SharedInformerFactory

	hpaVersion string
	hpaSii     cache.SharedIndexInformer
	hpaReg     cache.ResourceEventHandlerRegistration
}

func NewClient(hpaEventHandler cache.ResourceEventHandler, kubeConfig string) (*Client, error) {
	c := &Client{}

	if err := c.initClientSet(kubeConfig); err != nil {
		return nil, fmt.Errorf("[kubernetes] initClientSet error : %w", err)
	}
	c.iFactory = informers.NewSharedInformerFactory(c.cs, 0)

	if err := c.initShardIndexInformer(hpaEventHandler); err != nil {
		return nil, fmt.Errorf("[kubernetes] initShardIndexInformer error : %w", err)
	}

	return c, nil
}

func (c *Client) GetHPAVersion() string {
	return c.hpaVersion
}

func (c *Client) Start() {
	c.shutdown = make(chan struct{})
	c.iFactory.Start(c.shutdown)
}

func (c *Client) Stop() {
	_ = c.hpaSii.RemoveEventHandler(c.hpaReg)
	close(c.shutdown)
}

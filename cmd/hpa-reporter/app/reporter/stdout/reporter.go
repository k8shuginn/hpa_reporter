package stdout

import (
	"fmt"
	"github.com/k8shuginn/hpa_reporter/cmd/hpa-reporter/app/config"
	"github.com/k8shuginn/hpa_reporter/cmd/hpa-reporter/app/message"
)

// Reporter is stdout reporter
type Reporter struct {
	msgChan  chan *message.Data
	shutdown chan struct{}

	name    string
	configs map[string]string
}

// Report sends message to stdout
func (r *Reporter) Report(msg *message.Data) {
	r.msgChan <- msg
}

// run starts the reporter
func (r *Reporter) run() {
LOOP:
	for {
		select {
		case msg := <-r.msgChan:
			fmt.Printf("stdout(%s): %s[%s/%s]: replicas(%d/%d)\n", r.name, msg.Level, msg.Name, msg.Namespace, msg.CurrentReplicas, msg.MaxReplicas)
		case <-r.shutdown:
			break LOOP
		}
	}
}

// CreateReporter creates a new stdout reporter
func CreateReporter(cfg config.Reporter, shutdown chan struct{}) (*Reporter, error) {
	r := &Reporter{
		msgChan:  make(chan *message.Data),
		shutdown: shutdown,
		name:     cfg.Name,
		configs:  cfg.Configs,
	}
	go r.run()

	return r, nil
}

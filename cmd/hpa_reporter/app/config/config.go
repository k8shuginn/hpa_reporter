package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

const (
	DefaultConfigPath = "/etc/k8shuginn/config.yaml"
)

type (
	Reporter struct {
		Name    string            `yaml:"name"`
		Configs map[string]string `yaml:"configs"`
	}

	ReporterConfig struct {
		Slack  []Reporter `yaml:"slack"`
		Stdout []Reporter `yaml:"stdout"`
	}
)

type (
	HpaConfig struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
		Threshold int32  `yaml:"threshold"`
	}
)

type AppConfig struct {
	Reporters ReporterConfig `yaml:"reporters"`
	Hpa       []HpaConfig    `yaml:"hpa"`
}

// LoadConfig reads the configuration file and returns the AppConfig object
func LoadConfig(path string) (*AppConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var result AppConfig
	if err = yaml.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

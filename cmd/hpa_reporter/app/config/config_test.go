package config

import "testing"

func TestLoadConfig(t *testing.T) {
	path := "/Users/hsw/0_develop/0_code/k8shuginn/hpa_reporter/test/config.yml"
	cfg, err := LoadConfig(path)
	if err != nil {
		t.Errorf("failed to load config: %v", err)
	}

	t.Logf("config loaded: %v", cfg)
}

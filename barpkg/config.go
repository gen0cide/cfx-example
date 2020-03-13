package barpkg

import (
	"fmt"

	"github.com/gen0cide/cfx"
)

// NewBarConfig is a constructor that users the cfx package to inject the configuration from CFX parsed YAML.
func NewBarConfig(provider cfx.Container) (*Config, error) {
	// create an empty Config object
	cfg := &Config{}

	// use the provider to populate the config
	err := provider.Populate("bar", cfg)
	if err != nil {
		return nil, fmt.Errorf("error parsing bar config: %v", err)
	}

	// check for any errors in your config
	if cfg.Driver != "postgres" {
		return nil, fmt.Errorf("config value bar.driver does not equal 'postgres'")
	}

	return cfg, nil
}

// Config defines an example barpkg Config object
// (suppose to look like a database configuration)
type Config struct {
	Driver  string `json:"driver,omitempty" yaml:"driver,omitempty" mapstructure:"driver,omitempty"`
	ConnURL string `json:"conn_url,omitempty" yaml:"conn_url,omitempty" mapstructure:"conn_url,omitempty"`
}

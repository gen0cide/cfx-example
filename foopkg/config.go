package foopkg

import (
	"fmt"

	"github.com/gen0cide/cfx"
)

// NewFooConfig is a constructor that uses the cfx package to inject the configuration
// from CFX parsed YAML.
func NewFooConfig(provider cfx.Container) (*Config, error) {
	// create an empty Config object
	cfg := &Config{}

	// set defaults here if you want
	cfg.Age = 99

	// use the provider to populate the config
	err := provider.Populate("foo", cfg)
	if err != nil {
		return nil, fmt.Errorf("error parsing foo config: %v", err)
	}

	return cfg, nil
}

// Config defines an example foopkg Config object
type Config struct {
	Name string   `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`
	URL  string   `json:"url,omitempty" yaml:"url,omitempty" mapstructure:"url,omitempty"`
	Age  int      `json:"age,omitempty" yaml:"age,omitempty" mapstructure:"age,omitempty"`
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty" mapstructure:"tags,omitempty"`
}

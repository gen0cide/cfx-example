package barpkg

import (
	"fmt"

	"github.com/gen0cide/cfx"
)

// NewBar is a constructor that requires a populated config adn returns a Foo.
func NewBar(cfg *Config, env cfx.EnvContext) *Bar {
	return &Bar{
		Config: cfg,
		Env:    env,
	}
}

// Bar is an example container.
type Bar struct {
	Config *Config
	Env    cfx.EnvContext
}

// PrintInfo is used to show the features of the config object.
func (b *Bar) PrintInfo() {
	fmt.Printf("[BAR] Driver = %v\n", b.Config.Driver)
	fmt.Printf("[BAR] ConnURL = %v\n", b.Config.ConnURL)
}

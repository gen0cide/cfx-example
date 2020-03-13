package barpkg

import "go.uber.org/fx"

// Module defines the registered Fx constructors that can be used to resolve dependencies.
var Module = fx.Provide(
	NewBarConfig,
	NewBar,
)

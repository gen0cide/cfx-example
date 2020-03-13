package foopkg

import "go.uber.org/fx"

// Module defines the Fx constructors that can be used.
var Module = fx.Provide(
	NewFooConfig,
	NewFoo,
)

package middleware

import (
	"github.com/google/wire"
)

var MiddlewareWireModuleSet = wire.NewSet(
	MiddlewareWireSet,
)

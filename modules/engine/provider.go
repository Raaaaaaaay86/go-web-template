package engine

import (
	"github.com/google/wire"
)

var GinManagerModuleWireSet = wire.NewSet(
	GinManagerWireSet,
)

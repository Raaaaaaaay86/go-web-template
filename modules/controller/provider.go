package controller

import (
	"github.com/google/wire"
)

// All Controllers
var ControllerWireModuleSet = wire.NewSet(
	userControllerWireSet,
	contentControllerSet,
	rabbitMQControllerSet,
)

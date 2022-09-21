package service

import (
	"github.com/google/wire"
)

// All Services
var ServiceWireModuleSet = wire.NewSet(
	userServiceSet,
	contentServiceSet,
	rabbitMQServiceSet,
)

package repository

import (
	"github.com/google/wire"
)

var RepositoryWireModuleSet = wire.NewSet(
	userRepositorySet,
)

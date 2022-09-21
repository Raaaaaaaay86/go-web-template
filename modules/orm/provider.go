package orm

import (
	"go-web-template/modules/orm/mysql"

	"github.com/google/wire"
)

var OrmWireModuleSet = wire.NewSet(
	mysql.MySQLOrmSet,
)

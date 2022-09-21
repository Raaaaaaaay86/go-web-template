package orm

import (
	"go-web-template/modules/orm/mysql"

	"github.com/google/wire"
)

var OrmWireSet = wire.NewSet(
	mysql.MySQLOrmSet,
)
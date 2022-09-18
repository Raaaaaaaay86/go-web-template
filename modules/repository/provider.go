package repository

import "github.com/google/wire"

// * MySQL
var MySQLOrmSet = wire.NewSet(
	wire.Bind(new(IMySQLGorm), new(*MySQLGorm)),
	MySQLGormProvider,
)

func MySQLGormProvider() *MySQLGorm {
	return &MySQLGorm{}
}

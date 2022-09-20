package mysql

import (
	"os"

	"github.com/google/wire"
)

// * MySQL
var MySQLOrmSet = wire.NewSet(
	wire.Bind(new(IMySQLGorm), new(*MySQLGorm)),
	MySQLGormProvider,
)

func MySQLGormProvider() *MySQLGorm {
	return &MySQLGorm{
		DB_MYSQL_USERNAME: os.Getenv("DB_MYSQL_USERNAME"),
		DB_MYSQL_PASSWORD: os.Getenv("DB_MYSQL_PASSWORD"),
		DB_MYSQL_HOST:     os.Getenv("DB_MYSQL_HOST"),
		DB_MYSQL_PORT:     os.Getenv("DB_MYSQL_PORT"),
		DB_MYSQL_SCHEMA:   os.Getenv("DB_MYSQL_SCHEMA"),
	}
}

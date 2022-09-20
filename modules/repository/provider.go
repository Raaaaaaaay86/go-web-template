package repository

import (
	"go-web-template/modules/orm/mysql"

	"github.com/google/wire"
)

var RepositorySet = wire.NewSet(
	userRepositorySet,
)

// UserRepository
var userRepositorySet = wire.NewSet(
	wire.Bind(new(IUserRepository), new(UserRepository)),
	UserRepositoryProvider,
)

func UserRepositoryProvider(mysqlGorm mysql.IMySQLGorm) UserRepository {
	return UserRepository{
		MySQLGorm: mysqlGorm,
	}
}

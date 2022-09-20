package service

import (
	"go-web-template/modules/rabbitmq"
	"go-web-template/modules/orm/mysql"
	"go-web-template/modules/util/check"
	"go-web-template/modules/util/crypt"
	"go-web-template/modules/util/jwt"

	"github.com/google/wire"
)

// All Services
var ServiceSet = wire.NewSet(
	userServiceSet,
	contentServiceSet,
	rabbitMQServiceSet,
)

// UserService
var userServiceSet = wire.NewSet(
	wire.Bind(new(IUserService), new(UserService)),
	UserServiceProvider,
)

func UserServiceProvider(
	mySQLGorm *mysql.MySQLGorm,
	cryptTool crypt.PasswordCrypt,
	jwtManager jwt.JwtManager,
) UserService {
	return UserService{
		MySQLGorm:  mySQLGorm,
		CryptTool:  cryptTool,
		JwtManager: jwtManager,
	}
}

// ContentService
var contentServiceSet = wire.NewSet(
	wire.Bind(new(IContentService), new(ContentService)),
	ContentServiceProvider,
)

func ContentServiceProvider() ContentService {
	return ContentService{}
}

// RabbitMQService
var rabbitMQServiceSet = wire.NewSet(
	wire.Bind(new(IRabbitMQService), new(RabbitMQService)),
	RabbitMQServiceProvider,
)

func RabbitMQServiceProvider(rabbitMQManager rabbitmq.IRabbitMQManager, checker check.Checker) RabbitMQService {
	return RabbitMQService{
		RabbitMQManager: rabbitMQManager,
		Checker: checker,
	}
}

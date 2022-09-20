package engine

import (
	"go-web-template/modules/controller"
	"go-web-template/modules/orm/mysql"
	"go-web-template/modules/middleware"
	"go-web-template/modules/rabbitmq"
	"go-web-template/modules/service"
	"go-web-template/modules/util/check"
	"go-web-template/modules/util/crypt"
	"go-web-template/modules/util/jwt"

	"github.com/google/wire"
)

var GinManagerSet = wire.NewSet(
	wire.Bind(new(IGinManager), new(*GinManager)),
	GinManagerProvider,
	controller.ControllerWireSet,
	service.ServiceSet,
	mysql.MySQLOrmSet,
	crypt.PasswordCryptSet,
	jwt.JwtManagerSet,
	middleware.MiddlewareSet,
	rabbitmq.RabbitMQSet,
	check.CheckerProvider,
)

func GinManagerProvider(
	mysqlOrm *mysql.MySQLGorm,
	userController controller.IUserController,
	contentController controller.IContentController,
	rabbitMQController controller.IRabbitMQController,
	middleware middleware.IMiddleware,
	rabbitmqManager rabbitmq.IRabbitMQManager,
) *GinManager {
	return &GinManager{
		MySQLGorm:          mysqlOrm,
		Middleware:         middleware,
		UserController:     userController,
		ContentController:  contentController,
		RabbitMQController: rabbitMQController,
		RabbitMQManager:    rabbitMQController,
	}
}

package engine

import (
	"go-web-template/modules/controller"
	"go-web-template/modules/middleware"
	"go-web-template/modules/orm"
	"go-web-template/modules/orm/mysql"
	"go-web-template/modules/rabbitmq"
	"go-web-template/modules/repository"
	"go-web-template/modules/service"
	"go-web-template/modules/util"

	"github.com/google/wire"
)

var GinManagerSet = wire.NewSet(
	wire.Bind(new(IGinManager), new(*GinManager)),
	GinManagerProvider,
	controller.ControllerWireSet,
	service.ServiceSet,
	orm.OrmWireSet,
	util.UtilWireSet,
	rabbitmq.RabbitMQSet,
	repository.RepositorySet,
	middleware.MiddlewareSet,
)

func GinManagerProvider(
	mysqlOrm *mysql.MySQLGorm,
	middleware middleware.IMiddleware,
	userController controller.IUserController,
	contentController controller.IContentController,
	rabbitMQController controller.IRabbitMQController,
) *GinManager {
	return &GinManager{
		MySQLGorm:          mysqlOrm,
		Middleware:         middleware,
		UserController:     userController,
		ContentController:  contentController,
		RabbitMQController: rabbitMQController,
	}
}

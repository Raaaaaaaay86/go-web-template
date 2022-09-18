package engine

import (
	"go-web-template/modules/controller"
	"go-web-template/modules/middleware"
	"go-web-template/modules/repository"
	"go-web-template/modules/service"
	"go-web-template/modules/util/crypt"
	"go-web-template/modules/util/jwt"

	"github.com/google/wire"
)

var GinManagerSet = wire.NewSet(
	wire.Bind(new(IGinManager), new(*GinManager)),
	GinManagerProvider,
	controller.ControllerWireSet,
	service.ServiceSet,
	repository.MySQLOrmSet,
	crypt.PasswordCryptSet,
	jwt.JwtManagerSet,
	middleware.MiddlewareSet,
)

func GinManagerProvider(
	mysqlOrm *repository.MySQLGorm,
	userController controller.IUserController,
	contentController controller.IContentController,
	middleware middleware.IMiddleware,
) *GinManager {
	return &GinManager{
		MySQLGorm:         mysqlOrm,
		Middleware:        middleware,
		UserController:    userController,
		ContentController: contentController,
	}
}

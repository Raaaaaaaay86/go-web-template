package engine

import (
	_ "go-web-template/docs"
	"go-web-template/modules/controller"
	"go-web-template/modules/middleware"
	"go-web-template/modules/orm/mysql"

	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type IGinManager interface{}

var instance *gin.Engine

type GinManager struct {
	MySQLGorm          *mysql.MySQLGorm
	UserController     controller.IUserController
	ContentController  controller.IContentController
	RabbitMQController controller.IRabbitMQController
	Middleware         middleware.IMiddleware
}

var GinManagerWireSet = wire.NewSet(
	wire.Bind(new(IGinManager), new(*GinManager)),
	GinManagerProvider,
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

func (g *GinManager) GetGinEngine() *gin.Engine {
	if instance == nil {
		return g.createGin()
	}

	return instance
}

func (g *GinManager) createGin() *gin.Engine {
	if instance == nil {
		instance = gin.Default()
	}

	g.MySQLGorm.CreateMySQLConnection()

	instance.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRoute := instance.Group("/user")
	{
		userRoute.POST("/login", g.UserController.Login)
		userRoute.POST("/logout", g.UserController.Logout)
		userRoute.POST("/register", g.UserController.Register)
		userRoute.POST("/verify", g.UserController.Verify)
	}

	contentRoute := instance.Group("/content", g.Middleware.Verify)
	{
		contentRoute.GET("/random", g.ContentController.RandomContent)
	}

	rabbitMQRoute := instance.Group("/rabbitmq")
	{
		rabbitMQRoute.POST("/sendMessage", g.RabbitMQController.SendMessage)
	}

	return instance
}

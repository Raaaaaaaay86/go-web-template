package engine

import (
	_ "go-web-template/docs"
	"go-web-template/modules/controller"
	"go-web-template/modules/middleware"
	"go-web-template/modules/repository"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type IGinManager interface{}

var instance *gin.Engine

type GinManager struct {
	MySQLGorm         *repository.MySQLGorm
	UserController    controller.IUserController
	ContentController controller.IContentController
	Middleware        middleware.IMiddleware
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

	return instance
}

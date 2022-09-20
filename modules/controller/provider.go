package controller

import (
	"go-web-template/modules/service"

	"github.com/google/wire"
)

// All Controllers
var ControllerWireSet = wire.NewSet(
	userControllerWireSet,
	contentControllerSet,
	rabbitMQControllerSet,
)

// UserController
var userControllerWireSet = wire.NewSet(
	wire.Bind(new(IUserController), new(UserController)),
	UserControllerProvider,
)

func UserControllerProvider(userService service.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

// ContentController
var contentControllerSet = wire.NewSet(
	wire.Bind(new(IContentController), new(ContentController)),
	ContentControllerProvider,
)

func ContentControllerProvider(contentService service.IContentService) ContentController {
	return ContentController{
		ContentService: contentService,
	}
}

// RabbitMQController
var rabbitMQControllerSet = wire.NewSet(
	wire.Bind(new(IRabbitMQController), new(RabbitMQController)),
	RabbitMQControllerProvider,
)

func RabbitMQControllerProvider(rabbitMQService service.IRabbitMQService) RabbitMQController {
	return RabbitMQController{
		RabbitMQService: rabbitMQService,
	}
}

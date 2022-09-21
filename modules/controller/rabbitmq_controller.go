package controller

import (
	"go-web-template/modules/constant/exception"
	"go-web-template/modules/dto"
	"go-web-template/modules/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type IRabbitMQController interface {
	SendMessage(ctx *gin.Context)
}

type RabbitMQController struct {
	RabbitMQService service.IRabbitMQService
}

var rabbitMQControllerSet = wire.NewSet(
	wire.Bind(new(IRabbitMQController), new(RabbitMQController)),
	RabbitMQControllerProvider,
)

func RabbitMQControllerProvider(rabbitMQService service.IRabbitMQService) RabbitMQController {
	return RabbitMQController{
		RabbitMQService: rabbitMQService,
	}
}

// RandomContent godoc
// @Summary      Send RabbitMQ message to topic
// @Description  This message sending mode is "topic"
// @Tags         RabbitMQ
// @Param sendMessageData body JSONRequest[dto.SendMessageData] true "Message content"
// @Accept       json
// @Produce      plain
// @Router       /rabbitmq/sendMessage [post]
func (rc RabbitMQController) SendMessage(ctx *gin.Context) {
	var jsonRequest JSONRequest[dto.SendMessageData]

	if err := ctx.BindJSON(&jsonRequest); err != nil {
		log.Println(err.Error())
		handleError(ctx, http.StatusBadRequest, err)

		return
	}

	err := rc.RabbitMQService.SendMessage(jsonRequest.Data.Topic, jsonRequest.Data.Message)
	if err != nil {
		log.Println(err.Error())

		switch err {
		case exception.ErrInvalidData:
			handleError(ctx, http.StatusBadRequest, err)
		default:
			handleError(ctx, http.StatusInternalServerError, err)
		}

		return
	}

	handleOK(ctx, nil)
}

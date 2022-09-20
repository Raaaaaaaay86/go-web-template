package controller

import (
	"go-web-template/modules/dto"
	"go-web-template/modules/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IRabbitMQController interface {
	SendMessage(ctx *gin.Context)
}

type RabbitMQController struct {
	RabbitMQService service.IRabbitMQService
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
		handleError(ctx, http.StatusInternalServerError, err)

		return
	}

	handleOK(ctx, nil)
}

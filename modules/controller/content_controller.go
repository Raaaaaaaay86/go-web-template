package controller

import (
	"go-web-template/modules/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IContentController interface {
    RandomContent(ctx *gin.Context)
}

type ContentController struct {
    ContentService service.IContentService
}

func (cc ContentController) RandomContent(ctx *gin.Context) {
	ctx.String(
		http.StatusOK,
		"%s",
		cc.ContentService.RandomContent(),
	)
}

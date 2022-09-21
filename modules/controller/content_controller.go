package controller

import (
	"go-web-template/modules/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type IContentController interface {
	RandomContent(ctx *gin.Context)
}

type ContentController struct {
	ContentService service.IContentService
}

var contentControllerSet = wire.NewSet(
	wire.Bind(new(IContentController), new(ContentController)),
	ContentControllerProvider,
)

func ContentControllerProvider(contentService service.IContentService) ContentController {
	return ContentController{
		ContentService: contentService,
	}
}

// RandomContent godoc
// @Summary      Get Random UUID
// @Description  Before calling the API, you needs to set the JWT in the ```Authorization``` header.
// @Tags         Secured API
// @Accept       json
// @Produce      plain
// @Security     ApiKeyAuth
// @Router       /content/random [get]
func (cc ContentController) RandomContent(ctx *gin.Context) {
	ctx.String(
		http.StatusOK,
		"%s",
		cc.ContentService.RandomContent(),
	)
}

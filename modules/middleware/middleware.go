package middleware

import (
	"go-web-template/modules/util/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IMiddleware interface {
	Verify(ctx *gin.Context)
}

type Middleware struct {
	JwtManager jwt.IJwtManager
}

func (m Middleware) Verify(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	acceptTokenHead := "Bearer "
	tokenHead := token[0:len(acceptTokenHead)]

	if tokenHead != acceptTokenHead {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err := m.JwtManager.Verify(token[len(acceptTokenHead):])
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Next()
}

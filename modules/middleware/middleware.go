package middleware

import (
	"go-web-template/modules/util/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type IMiddleware interface {
	Verify(ctx *gin.Context)
}

type Middleware struct {
	JwtManager jwt.IJwtManager
}

var MiddlewareWireSet = wire.NewSet(
	wire.Bind(new(IMiddleware), new(Middleware)),
	MiddlewareProvider,
)

func MiddlewareProvider(jwtManager jwt.IJwtManager) Middleware {
	return Middleware{
		JwtManager: jwtManager,
	}
}

func (m Middleware) Verify(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	acceptTokenHead := "Bearer "

	if len(token) == 0 || len(token) <= len(acceptTokenHead) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

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

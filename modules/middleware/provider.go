package middleware

import (
	"go-web-template/modules/util/jwt"

	"github.com/google/wire"
)

var MiddlewareSet = wire.NewSet(
	wire.Bind(new(IMiddleware), new(Middleware)),
	MiddlewareProvider,
)

func MiddlewareProvider(jwtManager jwt.IJwtManager) Middleware {
	return Middleware{
		JwtManager: jwtManager,
	}
}

package jwt

import "github.com/google/wire"

var JwtManagerWireSet = wire.NewSet(
	wire.Bind(new(IJwtManager), new(JwtManager)),
	JwtManagerProvider,
)

func JwtManagerProvider() JwtManager {
	jwtManager := JwtManager{}
	jwtManager.setSigningKey("ddgdshkafgayaikshyvaksdghvdkasvgasdkyfgaysdjfggdshkafgayaikshyvaksdghvdkasvgasdkyfgaysdjfg")
	return jwtManager
}

package util

import (
	"go-web-template/modules/util/check"
	"go-web-template/modules/util/crypt"
	"go-web-template/modules/util/jwt"

	"github.com/google/wire"
)

var UtilWireSet = wire.NewSet(
	jwt.JwtManagerWireSet,
	crypt.PasswordCryptWireSet,
	check.CheckerWireSet,
)

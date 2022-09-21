package crypt

import "github.com/google/wire"

var PasswordCryptWireSet = wire.NewSet(
	wire.Bind(new(IPasswordCrypt), new(PasswordCrypt)),
	PasswordCryptProvider,
)

func PasswordCryptProvider() PasswordCrypt {
	return PasswordCrypt{}
}

package crypt

import (
	"github.com/google/wire"
	"golang.org/x/crypto/bcrypt"
)

//go:generate mockery --dir . --filename mock_crypt.go --name IPasswordCrypt --output ../../mocks
type IPasswordCrypt interface {
	Encode(password string) (encoded string, err error)
	Verify(hashedPassword, password string) (err error)
}

type PasswordCrypt struct{}

var PasswordCryptWireSet = wire.NewSet(
	wire.Bind(new(IPasswordCrypt), new(PasswordCrypt)),
	PasswordCryptProvider,
)

func PasswordCryptProvider() PasswordCrypt {
	return PasswordCrypt{}
}

func (pc PasswordCrypt) Encode(password string) (encoded string, err error) {
	encodedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return "", err
	}

	return string(encodedBytes), nil
}

func (pc PasswordCrypt) Verify(hashedPassword, password string) (err error) {
	err = bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
	if err != nil {
		return err
	}

	return nil
}

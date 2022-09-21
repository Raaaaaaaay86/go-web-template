package crypt

import (
	"golang.org/x/crypto/bcrypt"
)

type IPasswordCrypt interface {
	Encode(password string) (encoded string, err error)
	Verify(hashedPassword, password string) (err error)
}

type PasswordCrypt struct{}

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

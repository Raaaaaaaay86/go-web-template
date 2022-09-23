package test

import (
	"errors"
	"go-web-template/modules/constant/exception"
	"go-web-template/modules/mocks"
	"go-web-template/modules/model"
	"go-web-template/modules/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	registerEmail := "newUser@gmail.com"
	registerPassword := "strongPassword"
	hashedPassword := "anyPossibleHash"
	expectedToken := "thisIsValidToken"

	mockPasswordCrypt := mocks.NewIPasswordCrypt(t)
	mockJwtManager := mocks.NewIJwtManager(t)
	mockUserRepository := mocks.NewIUserRepository(t)
	svc := service.UserService{
		CryptTool:      mockPasswordCrypt,
		JwtManager:     mockJwtManager,
		UserRepository: mockUserRepository,
	}

	hashedPassword := "anyPossibleHash"
	mockUserRepository.On("FindByEmail", registerEmail).
		Return(
			model.User{
				Password: hashedPassword,
				Email:    registerEmail,
			},
			nil,
		)

	mockPasswordCrypt.On("Verify", hashedPassword, registerPassword).
		Return(nil)

	mockJwtManager.On("Create").
		Return(generateToken, nil)

	token, err := svc.Login(registerEmail, registerPassword)
	if err != nil {
	}

	fmt.Println(token)
}

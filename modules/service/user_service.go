package service

import (
	"go-web-template/modules/constant/exception"
	"go-web-template/modules/constant/role"
	"go-web-template/modules/dto"
	"go-web-template/modules/model"
	"go-web-template/modules/repository"
	"go-web-template/modules/util/crypt"
	"go-web-template/modules/util/jwt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type IUserService interface {
	Login(email, password string) (token string, err error)
	Logout(ctx *gin.Context)
	Register(registerData dto.RegisterData) (model.User, error)
	Verify(token string) error
}

type UserService struct {
	CryptTool      crypt.IPasswordCrypt
	JwtManager     jwt.IJwtManager
	UserRepository repository.IUserRepository
}

var userServiceSet = wire.NewSet(
	wire.Bind(new(IUserService), new(UserService)),
	UserServiceProvider,
)

func UserServiceProvider(
	cryptTool crypt.PasswordCrypt,
	jwtManager jwt.JwtManager,
	userRepository repository.UserRepository,
) UserService {
	return UserService{
		CryptTool:      cryptTool,
		JwtManager:     jwtManager,
		UserRepository: userRepository,
	}
}

func (us UserService) Login(email, password string) (token string, err error) {
	existUser, err := us.UserRepository.FindByEmail(email)
	if err != nil {
		return "", err
	}

	err = us.CryptTool.Verify(existUser.Password, password)
	if err != nil {
		return "", err
	}

	token, err = us.JwtManager.Create()
	if err != nil {
		return "", err
	}

	return token, nil
}

func (us UserService) Logout(ctx *gin.Context) {
	ctx.Request.Header.Set("Authorization", "")
}

func (us UserService) Register(registerData dto.RegisterData) (user model.User, err error) {
	if len(registerData.Email) == 0 || len(registerData.Password) == 0 {
		return user, exception.ErrInvalidEmailOrPassword
	}

	_, err = us.UserRepository.FindByEmail(registerData.Email)
	if err == nil {
		return user, exception.ErrEmailAlreadyTaken
	}

	hashedPassword, err := us.CryptTool.Encode(registerData.Password)
	if err != nil {
		log.Panicf("Unexpected error when hashing password: %s", err)
		return user, err
	}

	user = us.buildNormalUser(
		registerData.Email,
		hashedPassword,
		registerData.UserInfo,
	)

	err = us.UserRepository.Create(user)
	if err != nil {
		log.Printf("Create new User failed: %s\n", err)
		return user, exception.ErrRegisterFailed
	}

	return user, nil
}

func (us UserService) Verify(token string) error {
	acceptTokenHead := "Bearer "
	tokenType := token[0:len(acceptTokenHead)]

	if tokenType != acceptTokenHead {
		return exception.ErrInvalidJWT
	}

	err := us.JwtManager.Verify(token[len(acceptTokenHead):])
	if err != nil {
		return exception.ErrInvalidJWT
	}

	return nil
}

func (us UserService) buildNormalUser(email, password string, userInfo model.UserInfo) model.User {
	normalUserRole := model.UserRole{
		Name: role.USER,
		ID:   0,
	}

	return model.User{
		Email:    email,
		Password: password,
		UserRole: normalUserRole,
		UserInfo: userInfo,
	}
}

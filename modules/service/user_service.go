package service

import (
	"go-web-template/modules/constant/exception"
	"go-web-template/modules/constant/role"
	"go-web-template/modules/dto"
	"go-web-template/modules/model"
	"go-web-template/modules/orm/mysql"
	"go-web-template/modules/repository"
	"go-web-template/modules/util/crypt"
	"go-web-template/modules/util/jwt"
	"log"

	"github.com/gin-gonic/gin"
)

type IUserService interface {
	Login(email, password string) (token string, err error)
	Logout(ctx *gin.Context)
	Register(registerData dto.RegisterData) (model.User, error)
	Verify(token string) error
}

type UserService struct {
	MySQLGorm      *mysql.MySQLGorm
	CryptTool      crypt.PasswordCrypt
	JwtManager     jwt.JwtManager
	UserRepository repository.UserRepository
}

func (us UserService) Login(email, password string) (token string, err error) {
	db := us.MySQLGorm.Get()

	var existUser model.User

	tx := db.Where("email = ?", email).First(&existUser)
	if tx.Error != nil {
		return "", tx.Error
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

	_, tx := us.UserRepository.FindByEmail(registerData.Email)
	if tx.RowsAffected > 0 {
		return user, exception.ErrEmailAlreadyTaken
	}

	hashedPassword, err := us.CryptTool.Encode(registerData.Password)
	if err != nil {
		log.Panicf("Unexpected error when hashing password: %s", err)
		return user, err
	}

	user = model.User{
		Email:    registerData.Email,
		Password: hashedPassword,
		UserRole: model.UserRole{
			Name: role.USER,
			ID:   0,
		},
		UserInfo: registerData.UserInfo,
	}

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

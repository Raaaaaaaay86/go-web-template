package service

import (
    "go-web-template/modules/constant/exception"
    "go-web-template/modules/constant/role"
    "go-web-template/modules/model"
    "go-web-template/modules/repository"
    "go-web-template/modules/util/crypt"
    "go-web-template/modules/util/jwt"
    "log"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type IUserService interface {
    Login(email, password string) (token string, err error)
    Logout(ctx *gin.Context)
    Register(user model.User) (model.User, error)
    Verify(token string) error
}

type UserService struct {
    MySQLGorm  *repository.MySQLGorm
    CryptTool  crypt.PasswordCrypt
    JwtManager jwt.JwtManager
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

func (us UserService) Register(user model.User) (model.User, error) {
    if len(user.Email) == 0 || len(user.Password) == 0 {
        return user, exception.ErrInvalidEmailOrPassword
    }

    db := us.MySQLGorm.Get()

    var existCount int64

    db.Model(&user).Where("email = ?", user.Email).Count(&existCount)

    if existCount > 0 {
        log.Println("Email already taken")

        return user, exception.ErrEmailAlreadyTaken
    }

    err := db.Transaction(func(tx *gorm.DB) error {
        hashedPassword, err := us.CryptTool.Encode(user.Password)

        if err != nil {
            log.Panicln("Unexpected error when hashing password")

            return err
        }

        user.UserRoleId = 0
        user.UserRole.Name = role.USER
        user.Password = string(hashedPassword)

        if err := tx.Create(&user).Error; err != nil {
            log.Println("Create new account failed")

            return err
        }

        return nil
    })
    if err != nil {
        log.Printf("Register Failed: %s\n", err)

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

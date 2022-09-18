package controller

import (
    "go-web-template/modules/model"
    "go-web-template/modules/service"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
)

type IUserController interface {
    Login(ctx *gin.Context)
    Logout(ctx *gin.Context)
    Register(ctx *gin.Context)
    Verify(ctx *gin.Context)
}

type UserController struct {
    UserService service.UserService
}

func (uc UserController) Login(ctx *gin.Context) {
    var jsonData JSONRequest[model.User]

    if err := ctx.BindJSON(&jsonData); err != nil {
        log.Println(err.Error())
        handleError(ctx, http.StatusBadRequest, err)

        return
    }

    var loginUser = jsonData.Data

    token, err := uc.UserService.Login(loginUser.Email, loginUser.Password)
    if err != nil {
        log.Println(err.Error())
        handleError(ctx, http.StatusInternalServerError, err)

        return
    }

    handleOK(ctx, token)
}

func (uc UserController) Logout(ctx *gin.Context) {
    ctx.Request.Header.Set("Authorization", "")
    handleOK(ctx, nil)
}

func (uc UserController) Register(ctx *gin.Context) {
    var jsonData JSONRequest[model.User]

    if err := ctx.BindJSON(&jsonData); err != nil {
        log.Println(err.Error())
        handleError(ctx, http.StatusBadRequest, err)

        return
    }

    user, err := uc.UserService.Register(jsonData.Data)
    if err != nil {
        log.Println(err.Error())
        handleError(ctx, http.StatusInternalServerError, err)

        return
    }

    handleOK(ctx, user)
}

func (uc UserController) Verify(ctx *gin.Context) {
    token := ctx.Request.Header.Get("Authorization")

    err := uc.UserService.Verify(token)
    if err != nil {
        handleError(ctx, http.StatusUnauthorized, err)
        return
    }

    handleOK(ctx, nil)
}

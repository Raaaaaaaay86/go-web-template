package controller

import (
	"go-web-template/modules/dto"
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

// Login godoc
// @Summary      Login
// @Description  If login success, API will return the JWT in the response body
// @Tags         UserService
// @Param loginData body JSONRequest[dto.LoginData] true "Login email and password"
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /user/login [post]
func (uc UserController) Login(ctx *gin.Context) {
	var jsonData JSONRequest[dto.LoginData]

	if err := ctx.BindJSON(&jsonData); err != nil {
		log.Println(err.Error())
		handleError(ctx, http.StatusBadRequest, err)

		return
	}

	var loginData = jsonData.Data

	token, err := uc.UserService.Login(loginData.Email, loginData.Password)
	if err != nil {
		log.Println(err.Error())
		handleError(ctx, http.StatusInternalServerError, err)

		return
	}

	handleOK(ctx, token)
}

// Logout godoc
// @Summary      Logout
// @Description  Clear user's ```Authorization``` header
// @Tags         UserService
// @Accept       json
// @Produce      json
// @Router       /user/logout [post]
func (uc UserController) Logout(ctx *gin.Context) {
	ctx.Request.Response.Header.Set("Authorization", "")
	handleOK(ctx, nil)
}

// Register godoc
// @Summary      Register
// @Description  Register new user
// @Tags         UserService
// @Param registerUser body JSONRequest[dto.RegisterData] true "Register data"
// @Accept       json
// @Produce      json
// @Router       /user/register [post]
func (uc UserController) Register(ctx *gin.Context) {
	var jsonData JSONRequest[dto.RegisterData]

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

// Verify godoc
// @Summary      Verify
// @Description  Verify user JWT in ```Authorization``` header.
// @Tags         UserService
// @Accept       json
// @Produce      json
// @Router       /user/verify [get]
func (uc UserController) Verify(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	err := uc.UserService.Verify(token)
	if err != nil {
		handleError(ctx, http.StatusUnauthorized, err)
		return
	}

	handleOK(ctx, nil)
}

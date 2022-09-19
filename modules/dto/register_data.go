package dto

import "go-web-template/modules/model"

type RegisterData struct {
	Email    string         `json:"email"`
	Password string         `json:"password"`
	UserInfo model.UserInfo `json:"userInfo"`
}

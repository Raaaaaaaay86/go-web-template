package model

type User struct {
	ID         uint     `json:"-" gorm:"primaryKey"`
	Email      string   `json:"email"`
	Password   string   `json:"password"`
	UserRoleId uint     `json:"userRoleId"`
	UserRole   UserRole `json:"userRole"`
	UserInfo   UserInfo `json:"userInfo"`
}

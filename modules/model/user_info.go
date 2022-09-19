package model

type UserInfo struct {
	UserID    uint   `gorm:"primaryKey" json:"-"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
}

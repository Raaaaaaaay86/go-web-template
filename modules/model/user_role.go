package model

type UserRole struct {
    ID   uint   `gorm:"primaryKey" json:"-"`
    Name string `json:"name"`
}

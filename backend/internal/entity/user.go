package entity

import "gorm.io/gorm"

var UserTable = "users"

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}

type UserResetPassword struct {
	User
	NewPassword string `json:"newPassword"`
}

package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}

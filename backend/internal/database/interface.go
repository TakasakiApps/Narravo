package database

import (
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"gorm.io/gorm"
)

type Operation interface {
	Migrate(dst ...any)
}

type Driver interface {
	GetInstance() *gorm.DB
	Operation

	UserCtrl
}

// UserCtrl management for User
type UserCtrl interface {
	AddUser(user *entity.User) int64
	QueryUser(username string) *entity.User
}

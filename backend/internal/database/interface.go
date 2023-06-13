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
	// AddUser adds a new user to the database and returns the number of rows affected
	AddUser(user *entity.User) int64
	// QueryUser searches the database for a user with the given username and returns it if found
	QueryUser(username string) *entity.User
}

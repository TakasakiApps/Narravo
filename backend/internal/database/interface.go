package database

import "gorm.io/gorm"

type Operation interface {
	Migrate(dst ...any)
}

type Driver interface {
	GetDB() *gorm.DB
	Operation
}

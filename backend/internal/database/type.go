package database

import (
	"gorm.io/gorm"
)

type CommonDAO struct {
	Driver
	*gorm.DB
}

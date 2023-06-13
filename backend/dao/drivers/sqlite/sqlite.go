package sqlite

import (
	"github.com/TakasakiApps/Narravo/backend/internal/database"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/ohanakogo/exceptiongo"
	"gorm.io/gorm"
)

type SQLite database.CommonDAO

func (s SQLite) GetInstance() *gorm.DB {
	return s.DB
}

func (s SQLite) Migrate(dst ...any) {
	err := s.AutoMigrate(dst...)
	exceptiongo.QuickThrow[types.DataBaseOperationFailedException](err)
}

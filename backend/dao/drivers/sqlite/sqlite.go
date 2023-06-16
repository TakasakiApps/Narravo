package sqlite

import (
	"github.com/TakasakiApps/Narravo/backend/internal/database"
	"gorm.io/gorm"
)

type SQLite database.CommonDAO

func (s SQLite) GetInstance() *gorm.DB {
	return s.DB
}

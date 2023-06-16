package dao

import (
	"github.com/TakasakiApps/Narravo/backend/dao/drivers/sqlite"
	"github.com/TakasakiApps/Narravo/backend/internal/database"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/ohanakogo/exceptiongo"
	"github.com/ohanakogo/ohanakoutilgo"
	SQLiteDriver "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewDAO returns a DAO instance initialized with the provided URL and driver type
func NewDAO[T database.Driver](url string) (dao database.Driver) {
	db, err := gorm.Open(SQLiteDriver.Open(url), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	exceptiongo.QuickThrow[types.DataBaseConnectFailedException](err)

	switch ohanakoutilgo.TypeOf[T]() {
	case ohanakoutilgo.TypeOf[sqlite.SQLite]():
		dao = &sqlite.SQLite{DB: db}
		return
	}

	// Return nil if no valid driver type is found
	return nil
}

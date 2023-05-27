package dao

import (
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/dao/sqlite"
	"github.com/TakasakiApps/Narravo/backend/internal/database"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/ohanakogo/exceptiongo"
	"github.com/ohanakogo/ohanakoutilgo"
	SQLiteDriver "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewDAO returns a DAO instance initialized with the provided URL and driver type
func NewDAO[T database.Driver](url string) (dao database.Driver) {
	db, err := gorm.Open(SQLiteDriver.Open(url), &gorm.Config{})
	exceptiongo.QuickThrow[types.DataBaseConnectFailedException](err)

	switch ohanakoutilgo.TypeOf[T]() {
	case ohanakoutilgo.TypeOf[sqlite.SQLite]():
		dao = &sqlite.SQLite{DB: db}
		return
	default:
		// If the driver type is not supported, a runtime exception will be thrown
		exceptiongo.QuickThrowMsg[types.RuntimeException](
			fmt.Sprintf("Failed to instantiate DAO of type <%v> with the driver", ohanakoutilgo.TypeOf[T]()))
	}

	// Return nil if no valid driver type is found
	return nil
}

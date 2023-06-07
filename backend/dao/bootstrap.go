package dao

import (
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/dao/drivers/sqlite"
	"github.com/TakasakiApps/Narravo/backend/internal/database"
	"github.com/TakasakiApps/Narravo/backend/internal/global"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/ohanakogo/exceptiongo"
)

var driver database.Driver

func Get() database.Driver {
	return driver
}

type DataBaseBootstrapConfig struct {
	DriverType types.DataBaseDriverType
	Url        string
}

//goland:noinspection GoUnhandledErrorResult
func DataBaseBootstrap(c *DataBaseBootstrapConfig) {
	global.GetLogger().Infof("Bootstrapping DAO of type <%s>", c.DriverType)
	switch c.DriverType {
	case types.SQLITE:
		driver = NewDAO[sqlite.SQLite](c.Url)
	default:
		// If the driver type is not supported, a runtime exception will be thrown
		exceptiongo.QuickThrowMsg[types.RuntimeException](
			fmt.Sprintf("Failed to instantiate DAO of type <%s> with the driver", c.DriverType))
	}

	global.GetLogger().Info("Migrating database...")
	driver.Migrate()
}

package dao

import (
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/dao/drivers/sqlite"
	"github.com/TakasakiApps/Narravo/backend/internal/database"
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"github.com/TakasakiApps/Narravo/backend/internal/global"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/ohanakogo/exceptiongo"
)

var driver database.Driver

func GetInstance() database.Driver {
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

	migrateDataBase()

	initGuestUser()
}

// migrateDataBase to migrate structure database
func migrateDataBase() {
	global.GetLogger().Info("Start migrating database...")

	migrator := driver.GetInstance().Migrator()

	targets := []types.Pair[any, string]{
		{&entity.User{}, "users"},
		{&entity.BookLocalInfo{}, "book_infos"},
	}

	for _, obj := range targets {
		if !migrator.HasTable(obj.First) {
			global.GetLogger().Infof("Migrating dataBase<%s>", obj.Second)
			err := migrator.CreateTable(obj.First)
			exceptiongo.QuickThrow[types.DataBaseOperationFailedException](err)
			if !migrator.HasTable(obj.Second) {
				err = migrator.RenameTable(obj.First, obj.Second)
				exceptiongo.QuickThrow[types.DataBaseOperationFailedException](err)
			}
		}
	}
}

// initGuestUser to initialize the guest user
func initGuestUser() {
	// Create a guest user with default credentials
	guest := &entity.User{
		Name:     "guest",
		Password: "guest",
	}
	// Check if the guest user already exists in the database
	queryGuest := driver.QueryUserByName(guest)
	if queryGuest == nil {
		// If the guest user does not exist, add it to the database
		global.GetLogger().Info("Users: Guest user not found")
		driver.AddUser(guest)
		global.GetLogger().Infof("Users: Guest user created successfully, user<username: %s, password: %s>", guest.Name, guest.Password)
	}
}

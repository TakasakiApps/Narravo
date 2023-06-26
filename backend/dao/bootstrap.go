package dao

import (
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/dao/drivers/sqlite"
	"github.com/TakasakiApps/Narravo/backend/internal/database"
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"github.com/TakasakiApps/Narravo/backend/internal/global"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
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

	initUser(
		"admin",
		&entity.User{
			Name:     "admin",
			Password: utils.GenerateAnyID(12),
		},
	)
	initUser(
		"guest",
		&entity.User{
			Name:     "guest",
			Password: "guest",
		},
	)
}

// migrateDataBase to migrate structure database
func migrateDataBase() {
	global.GetLogger().Info("Start migrating database...")

	migrator := driver.GetInstance().Migrator()

	targets := []types.Pair[any, string]{
		{&entity.IdRecord{}, entity.IdRecordTable},
		{&entity.User{}, entity.UserTable},
		{&entity.Author{}, entity.AuthorTable},
		{&entity.Novel{}, entity.NovelTable},
	}

	for _, obj := range targets {
		if !migrator.HasTable(obj.First) && !migrator.HasTable(obj.Second) {
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

// initUser to initialize the guest user
func initUser(userType string, user *entity.User) {
	// Check if the guest user already exists in the database
	queryUser := driver.QueryUserByName(user.Name)
	if queryUser == nil {
		// If the guest user does not exist, add it to the database
		global.GetLogger().Infof("Users: %s user not found", userType)
		driver.AddUser(&entity.User{
			Name:     user.Name,
			Password: utils.MD5(user.Password),
		})
		global.GetLogger().Infof("Users: %s user created successfully, user<username: %s, password: %s>", userType, user.Name, user.Password)
	}
}

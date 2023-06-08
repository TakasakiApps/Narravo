package test

import (
	"encoding/json"
	"github.com/TakasakiApps/Narravo/backend/dao"
	"github.com/TakasakiApps/Narravo/backend/dao/drivers/sqlite"
	"github.com/TakasakiApps/Narravo/backend/internal/database"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/ohanakogo/exceptiongo"
	"testing"
)

var driver database.Driver

type Test struct {
	Id     int     `gorm:"primaryKey;autoIncrement"`
	Param1 string  `gorm:"not null"`
	Param2 int     `gorm:"not null"`
	Param3 float64 `gorm:"not null"`
}

func TestSQLite(t *testing.T) {
	driver = dao.NewDAO[sqlite.SQLite]("test.db")

	defer func() {
		db, _ := driver.GetDB().DB()
		_ = db.Close()
		fsutil.MustRemove("test.db")
	}()

	// Test migrate
	migrate()

	// add
	add()

	// remove
	remove()

	// update
	update()

	// query
	queryResult := query()

	marshal, err := json.Marshal(queryResult)
	exceptiongo.QuickThrow[types.JsonUnmarshalFailedException](err)
	assert.Eq(t, `{"Id":2,"Param1":"updated","Param2":1,"Param3":3.14}`, string(marshal))
}

func migrate() {
	driver.Migrate(&Test{})
}

func add() {
	for i := 0; i < 3; i++ {
		driver.GetDB().Model(&Test{}).Create(&Test{
			Param1: "test",
			Param2: 1 * i,
			Param3: 3.14 * float64(i),
		})
	}
}

func remove() {
	driver.GetDB().Model(&Test{}).Delete(&Test{Id: 1})
}

func update() {
	driver.GetDB().Model(&Test{}).
		Select("param1").
		Where("id = ?", 2).
		Updates(&Test{
			Param1: "updated",
		})
}

func query() Test {
	queryTarget := Test{Id: 2}
	driver.GetDB().Model(&Test{}).First(&queryTarget)
	return queryTarget
}

package bootstrap

import (
	"github.com/TakasakiApps/Narravo/backend/dao"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
)

func Start() {
	// hardcode database type and path
	dao.DataBaseBootstrap(&dao.DataBaseBootstrapConfig{
		DriverType: types.SQLITE,
		Url:        "data.db",
	})
}

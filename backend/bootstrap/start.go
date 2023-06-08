package bootstrap

import (
	"github.com/TakasakiApps/Narravo/backend/dao"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"path/filepath"
)

func Start() {
	// hardcode database type and path
	dao.DataBaseBootstrap(&dao.DataBaseBootstrapConfig{
		DriverType: types.SQLITE,
		Url:        filepath.Join(utils.EnsureDir(utils.GetAppHomeDir()), "data.db"),
	})
}

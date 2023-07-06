package bootstrap

import (
	"github.com/TakasakiApps/Narravo/backend/config"
	"github.com/TakasakiApps/Narravo/backend/dao"
	"github.com/TakasakiApps/Narravo/backend/internal/global"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/gookit/goutil/fsutil"
	"github.com/joho/godotenv"
	"path/filepath"
)

func bootstrapDotEnv() {
	if fsutil.FileExists(global.DotEnvFile) {
		err := godotenv.Load(global.DotEnvFile)
		global.GetLogger().Warn(err)
	}
}

func Start() {
	config.Bootstrap()

	bootstrapDotEnv()

	// hardcode database type and path
	dao.DataBaseBootstrap(&dao.DataBaseBootstrapConfig{
		DriverType: types.SQLITE,
		Url:        filepath.Join(utils.EnsureDir(utils.GetAppHomeDir()), "data.db"),
	})
}

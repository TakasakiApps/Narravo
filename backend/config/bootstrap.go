package config

import (
	"encoding/json"
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/internal/global"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/gookit/goutil/fsutil"
	"github.com/ohanakogo/exceptiongo"
	"os"
	"path/filepath"
)

// path contains the file path of the configuration file.
var path = filepath.Join(utils.GetAppHomeDir(), "config.json")

// Bootstrap initializes the configuration by checking if the config file exists. If it does not exist, a new one is created with default values;
// otherwise, its contents are read and deserialized into the application's configuration object.
func Bootstrap() {
	global.GetLogger().Info("Bootstrapping configuration")
	if !checkConfigFile() {
		global.GetLogger().Info("Configuration file not found, trying to initialize...")
		marshalIndent, _ := json.MarshalIndent(GetInstance(), "", "    ")
		err := fsutil.WriteFile(path, marshalIndent, os.ModePerm)
		exceptiongo.QuickThrow[types.IOException](err)
	} else {
		global.GetLogger().Info("Configuration file found, reading...")
		readConfig := fsutil.ReadFile(path)
		err := json.Unmarshal(readConfig, mConfig)
		if err != nil {
			global.GetLogger().Error("Failed to parse configuration, trying to delete configuration file and re-initialize program. Now exiting...")
			exceptiongo.QuickThrow[types.JsonUnmarshalFailedException](err)
		}
	}
}

// checkConfigFile checks if the configuration file exists and returns true if it does.
func checkConfigFile() bool {
	if fsutil.PathExists(path) {
		if fsutil.IsDir(path) {
			exceptiongo.QuickThrowMsg[types.FileSystemException](
				fmt.Sprintf("target config path %s is a directory", path),
			)
		}
		return true
	}
	return false
}

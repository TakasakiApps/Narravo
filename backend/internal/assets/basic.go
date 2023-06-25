package assets

import (
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/gookit/goutil/fsutil"
	"github.com/ohanakogo/exceptiongo"
	"path/filepath"
)

func ensureAssetsDir() string {
	return utils.EnsureDir(utils.GetAppHomeDir(), "assets")
}

func ensureAssetsSubDir(subDirs ...string) string {
	dir := filepath.Join(ensureAssetsDir(), filepath.Join(subDirs...))
	return utils.EnsureDir(dir)
}

func readBytes(path string) []byte {
	if fsutil.FileExists(path) {
		return fsutil.ReadFile(path)
	}
	return nil
}

func writeAny[T string | []byte](path string, data T) bool {
	if fsutil.FileExists(path) {
		err := fsutil.Remove(path)
		exceptiongo.QuickThrow[types.FileSystemException](err)
	}
	err := fsutil.WriteFile(path, data, 0755)
	exceptiongo.QuickThrow[types.FileSystemException](err)
	return true
}

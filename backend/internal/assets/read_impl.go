package assets

import (
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"path/filepath"
)

func ensureAssetsDir() string {
	return utils.EnsureDir(utils.GetAppHomeDir(), "assets")
}

func ensureAssetsSubDir(subDirs ...string) string {
	dir := filepath.Join(ensureAssetsDir(), filepath.Join(subDirs...))
	return utils.EnsureDir(dir)
}

func readFile(assetType AssetType, filename string) ([]byte, bool) {
	dir := ensureAssetsSubDir("files", string(assetType))
	path := filepath.Join(dir, filename)
	read := readBytes(path)
	if read == nil {
		return nil, false
	}
	return read, true
}

func readFileString(assetType AssetType, filename string) (string, bool) {
	bytes, exists := readFile(assetType, filename)
	if exists {
		return string(bytes), exists
	}
	return "", exists
}

func readImage(assetType AssetType, filename string) ([]byte, bool) {
	dir := ensureAssetsSubDir("images", string(assetType))
	path := filepath.Join(dir, filename)
	read := readBytes(path)
	if read == nil {
		return nil, false
	}
	return read, true
}

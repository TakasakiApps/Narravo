package assets

import (
	"github.com/gookit/goutil/fsutil"
	"path/filepath"
)

func deleteFile(assetType AssetType, filename string, fileType FileType) bool {
	dir := ensureAssetsSubDir(string(fileType), string(assetType))
	path := filepath.Join(dir, filename)
	err := fsutil.DeleteIfFileExist(path)
	if err != nil {
		return false
	}
	return true
}

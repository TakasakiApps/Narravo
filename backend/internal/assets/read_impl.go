package assets

import (
	"path/filepath"
)

func readFile(assetType AssetType, filename string, fileType FileType) ([]byte, bool) {
	dir := ensureAssetsSubDir(string(fileType), string(assetType))
	path := filepath.Join(dir, filename)
	read := readBytes(path)
	if read == nil {
		return nil, false
	}
	return read, true
}

func readFileString(assetType AssetType, filename string) (string, bool) {
	bytes, exists := readFile(assetType, filename, File)
	if exists {
		return string(bytes), exists
	}
	return "", exists
}

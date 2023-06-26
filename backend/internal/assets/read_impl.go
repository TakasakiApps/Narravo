package assets

import (
	"path/filepath"
)

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
	dir := ensureAssetsSubDir("images", string(assetType), filename)
	path := filepath.Join(dir, filename)
	read := readBytes(path)
	if read == nil {
		return nil, false
	}
	return read, true
}

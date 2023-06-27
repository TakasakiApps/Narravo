package assets

import "path/filepath"

func writeFile[T string | []byte](assetType AssetType, filename string, fileType FileType, data T) bool {
	path := getWriteFilePath(assetType, filename, fileType)
	return writeAny[T](path, data)
}

func getWriteFilePath(assetType AssetType, filename string, fileType FileType) string {
	dir := ensureAssetsSubDir(string(fileType), string(assetType), filepath.Dir(filename))
	path := filepath.Join(dir, filepath.Base(filename))
	return path
}

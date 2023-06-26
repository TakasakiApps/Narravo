package assets

import "path/filepath"

func writeFile[T string | []byte](assetType AssetType, filename string, fileType FileType, data T) bool {
	dir := ensureAssetsSubDir(string(fileType), string(assetType), filepath.Dir(filename))
	path := filepath.Join(dir, filepath.Base(filename))
	return writeAny[T](path, data)
}

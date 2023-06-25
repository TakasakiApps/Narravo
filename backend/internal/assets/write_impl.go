package assets

import "path/filepath"

func writeFile[T string | []byte](assetType AssetType, filename string, data T) bool {
	dir := ensureAssetsSubDir("files", string(assetType), filepath.Dir(filename))
	path := filepath.Join(dir, filepath.Base(filename))
	return writeAny[T](path, data)
}

func writeImage(assetType AssetType, filename string, data []byte) bool {
	dir := ensureAssetsSubDir("images", string(assetType), filepath.Dir(filename))
	path := filepath.Join(dir, filepath.Base(filename))
	return writeAny[[]byte](path, data)
}

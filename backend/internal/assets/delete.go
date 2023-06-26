package assets

import (
	"path/filepath"
)

func DeleteCover(coverId string) bool {
	return deleteFile(Cover, coverId, Image)
}

func DeleteAvatar(avatarId string) bool {
	return deleteFile(Avatar, avatarId, Image)
}

func DeleteCatalog(novelId string) bool {
	return deleteFile(Catalog, filepath.Join(novelId, "catalog.json"), File)
}

func DeleteChapter(novelId string, chapterId string) bool {
	return deleteFile(Catalog, filepath.Join(novelId, chapterId), File)
}

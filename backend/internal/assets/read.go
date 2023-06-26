package assets

import (
	"encoding/json"
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/ohanakogo/exceptiongo"
	"path/filepath"
)

func ReadCover(coverId string) ([]byte, bool) {
	return readFile(Cover, coverId, Image)
}

func ReadAvatar(avatarId string) ([]byte, bool) {
	return readFile(Avatar, avatarId, Image)
}

func ReadCatalog(novelId string) (*entity.CatalogInfo, bool) {
	bytes, exists := readFile(Catalog, filepath.Join(novelId, "catalog.json"), File)
	if !exists {
		return nil, exists
	}
	catalogInfo := entity.CatalogInfo{}
	err := json.Unmarshal(bytes, &catalogInfo)
	exceptiongo.QuickThrow[types.JsonUnmarshalFailedException](err)
	return &catalogInfo, exists
}

func ReadChapter(novelId string, chapterId string) (string, bool) {
	return readFileString(Catalog, filepath.Join(novelId, chapterId))
}

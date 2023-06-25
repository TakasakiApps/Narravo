package assets

import (
	"encoding/json"
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/ohanakogo/exceptiongo"
	"path/filepath"
)

func ReadCover(coverId string) ([]byte, bool) {
	return readImage(Cover, coverId)
}

func ReadAvatar(avatarId string) ([]byte, bool) {
	return readImage(Cover, avatarId)
}

func ReadCatalog(catalogId string) (*entity.CatalogInfo, bool) {
	bytes, exists := readFile(Catalog, filepath.Join(catalogId, "catalog.json"))
	if !exists {
		return nil, exists
	}
	catalogInfo := entity.CatalogInfo{}
	err := json.Unmarshal(bytes, &catalogInfo)
	exceptiongo.QuickThrow[types.JsonUnmarshalFailedException](err)
	return &catalogInfo, exists
}

func ReadChapter(catalogId string, chapterId string) (string, bool) {
	return readFileString(Catalog, filepath.Join(catalogId, chapterId))
}

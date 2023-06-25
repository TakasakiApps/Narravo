package assets

import (
	"encoding/json"
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/ohanakogo/exceptiongo"
)

func ReadCover(coverId string) ([]byte, bool) {
	return readImage(Cover, coverId)
}

func ReadAvatar(avatarId string) ([]byte, bool) {
	return readImage(Cover, avatarId)
}

func ReadCatalog(catalogId string) (*entity.CatalogInfo, bool) {
	bytes, exists := readFile(Catalog, catalogId+".json")
	if !exists {
		return nil, exists
	}
	catalogInfo := entity.CatalogInfo{}
	err := json.Unmarshal(bytes, &catalogInfo)
	exceptiongo.QuickThrow[types.JsonUnmarshalFailedException](err)
	return &catalogInfo, exists
}

func ReadChapter(chapterId string) (string, bool) {
	return readFileString(Catalog, chapterId)
}

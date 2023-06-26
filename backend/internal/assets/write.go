package assets

import (
	"encoding/json"
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/ohanakogo/exceptiongo"
	"path/filepath"
)

func WriteCover(coverId string, data []byte) bool {
	return writeImage(Cover, coverId, data)
}

func WriteAvatar(avatarId string, data []byte) bool {
	return writeImage(Cover, avatarId, data)
}

func WriteCatalog(novelId string, catalogInfo *entity.CatalogInfo) bool {
	data, err := json.MarshalIndent(catalogInfo, "", "    ")
	exceptiongo.QuickThrow[types.JsonMarshalFailedException](err)
	return writeFile[[]byte](Catalog, filepath.Join(novelId, "catalog.json"), data)
}

func WriteChapter(novelId string, chapterId string, data string) bool {
	return writeFile[string](Catalog, filepath.Join(novelId, chapterId), data)
}

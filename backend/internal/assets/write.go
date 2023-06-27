package assets

import (
	"encoding/json"
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/ohanakogo/exceptiongo"
	"path/filepath"
)

func GetCoverPath(coverId string) string {
	return getWriteFilePath(Cover, coverId, Image)
}

func GetAvatarPath(avatarId string) string {
	return getWriteFilePath(Avatar, avatarId, Image)
}

func WriteCatalog(novelId string, catalogInfo *entity.CatalogInfo) bool {
	data, err := json.MarshalIndent(catalogInfo, "", "    ")
	exceptiongo.QuickThrow[types.JsonMarshalFailedException](err)
	return writeFile[[]byte](Catalog, filepath.Join(novelId, "catalog.json"), File, data)
}

func WriteChapter(novelId string, chapterId string, data string) bool {
	return writeFile[string](Catalog, filepath.Join(novelId, chapterId), File, data)
}

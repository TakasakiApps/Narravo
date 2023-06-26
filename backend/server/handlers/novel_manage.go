package handlers

import (
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/dao"
	"github.com/TakasakiApps/Narravo/backend/internal/assets"
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/arrutil"
	"github.com/ohanakogo/exceptiongo"
	"github.com/ohanakogo/ohanakoutilgo"
	"net/http"
)

var UpdateNovelCatalog gin.HandlerFunc = func(c *gin.Context) {
	novelId := c.Param("novelId")

	novel := dao.GetInstance().QueryNovelById(novelId)
	if novel == nil {
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException](fmt.Sprintf("novel<%s> not found", novelId))
	}

	catalog := &entity.CatalogInfo{}
	utils.ConvMapToStructure(utils.GetData(c), catalog)

	localCatalog, exists := assets.ReadCatalog(novelId)
	if exists {
		diff := arrutil.Differences[entity.ChapterInfo](localCatalog.Chapters, catalog.Chapters, func(a, b any) int {
			chapterA := ohanakoutilgo.CastTo[entity.ChapterInfo](a)
			chapterB := ohanakoutilgo.CastTo[entity.ChapterInfo](b)

			if chapterA.ChapterId == "" || chapterB.ChapterId == "" || chapterA.ChapterId == chapterB.ChapterId {
				return 0
			}

			return -1
		})

		for _, chapterInfo := range diff {
			if _, exists := assets.ReadChapter(novelId, chapterInfo.ChapterId); !exists {
				continue
			}
			assets.DeleteChapter(novelId, chapterInfo.ChapterId)
		}
	}

	chapters := &catalog.Chapters
	for i := 0; i < len(*chapters); i++ {
		chapter := &(*chapters)[i]
		if chapter.ChapterId == "" {
			chapter.ChapterId = GenerateUniqueID()
			assets.WriteChapter(novelId, chapter.ChapterId, "")
		}
	}

	if assets.WriteCatalog(novelId, catalog) {
		c.JSON(http.StatusOK, catalog)
	} else {
		exceptiongo.QuickThrowMsg[types.ServerNotModifiedException]("update catalog failed")
	}
}

var AddNovel gin.HandlerFunc = func(c *gin.Context) {
	postNovel := &entity.PostNovel{}
	utils.ConvMapToStructure(utils.GetData(c), postNovel)

	_, exists := assets.ReadCover(postNovel.CoverId)
	if !exists {
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException]("cover not found, you must upload a cover before add novel")
	}

	author := dao.GetInstance().QueryAuthorById(postNovel.AuthorId)
	if author == nil {
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException]("author not found, you must select a valid author or create an author before add novel")
	}

	dao.GetInstance().AddNovel(&entity.Novel{
		ID:        GenerateUniqueID(),
		PostNovel: *postNovel,
	})
}

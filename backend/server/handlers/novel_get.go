package handlers

import (
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/dao"
	"github.com/TakasakiApps/Narravo/backend/internal/assets"
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/strutil"
	"github.com/ohanakogo/exceptiongo"
	"math"
	"net/http"
)

func toResponsibleNovelInfo(novel *entity.Novel) map[string]any {
	return map[string]any{
		"id":          novel.ID,
		"name":        novel.Name,
		"description": novel.Description,
		"author":      dao.GetInstance().QueryAuthorById(novel.AuthorId),
		"coverId":     novel.CoverId,
	}
}

func doPaginationNovelQuery(c *gin.Context, fn func(pageInt int, countInt int) []*entity.Novel) {
	page := c.Query("page")
	pageInt, err := strutil.Int(page)
	if err != nil {
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException](fmt.Sprintf("page<%v> isn't a int value", page))
	}
	count := c.Query("count")
	countInt, err := strutil.Int(count)
	if err != nil {
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException](fmt.Sprintf("count<%v> isn't a int value", count))
	}

	countNovel := dao.GetInstance().CountNovel()
	maxPage := int(math.Ceil(float64(countNovel) / float64(countInt)))

	var novelInfoList []map[string]any

	if pageInt <= maxPage {
		novels := fn(pageInt, countInt)
		for _, novel := range novels {
			novelInfoList = append(novelInfoList, toResponsibleNovelInfo(novel))
		}
	}

	c.JSON(http.StatusOK, map[string]any{
		"maxPage": maxPage,
		"list":    novelInfoList,
	})
}

var GetAllNovelInfo gin.HandlerFunc = func(c *gin.Context) {
	doPaginationNovelQuery(c, func(pageInt int, countInt int) []*entity.Novel {
		return dao.GetInstance().QueryNovelsLimit((pageInt-1)*countInt, countInt)
	})
}

var GetNovelInfo gin.HandlerFunc = func(c *gin.Context) {
	novelId := c.Param("novelId")

	novel := dao.GetInstance().QueryNovelById(novelId)
	if novel == nil {
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException](fmt.Sprintf("novel<%s> not found", novelId))
	}

	c.JSON(
		http.StatusOK,
		toResponsibleNovelInfo(novel),
	)
}

var SearchNovelInfo gin.HandlerFunc = func(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException](`query argument "keyword" can't be empty'`)
	}
	doPaginationNovelQuery(c, func(pageInt int, countInt int) []*entity.Novel {
		return dao.GetInstance().SearchNovelByName(keyword, (pageInt-1)*countInt, countInt)
	})
}

var GetNovelCatalog gin.HandlerFunc = func(c *gin.Context) {
	novelId := c.Param("novelId")

	catalog, exists := assets.ReadCatalog(novelId)
	if !exists {
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException](fmt.Sprintf("assets error: catalog not found"))
	}

	c.JSON(http.StatusOK, catalog)
}

var GetNovelChapter gin.HandlerFunc = func(c *gin.Context) {
	novelId := c.Param("novelId")
	chapterId := c.Query("chapterId")

	chapterContent, exists := assets.ReadChapter(novelId, chapterId)
	if !exists {
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException](fmt.Sprintf("assets error: chapter<%s> was not found in novel<%s>", chapterId, novelId))
	}

	c.String(http.StatusOK, chapterContent)
}

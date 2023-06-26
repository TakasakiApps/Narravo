package handlers

import (
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/dao"
	"github.com/TakasakiApps/Narravo/backend/internal/assets"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/ohanakogo/exceptiongo"
	"net/http"
)

var GetNovelInfo gin.HandlerFunc = func(c *gin.Context) {
	novelId := c.Param("novel_id")

	novel := dao.GetInstance().QueryNovelById(novelId)

	if novel == nil {
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException](fmt.Sprintf("novel<%s> not found", novelId))
	}

	c.JSON(
		http.StatusOK,
		map[string]string{
			"name":        novel.Name,
			"description": novel.Description,
			"authorId":    novel.AuthorId,
			"coverId":     novel.CoverId,
		},
	)
}

var GetNovelCatalog gin.HandlerFunc = func(c *gin.Context) {
	novelId := c.Param("novel_id")

	catalog, exists := assets.ReadCatalog(novelId)
	if !exists {
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException](fmt.Sprintf("assets error: catalog not found"))
	}

	c.JSON(http.StatusOK, catalog)
}

var GetNovelChapter gin.HandlerFunc = func(c *gin.Context) {
	novelId := c.Param("novel_id")
	chapterId := c.Query("chapterId")

	chapterContent, exists := assets.ReadChapter(novelId, chapterId)
	if !exists {
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException](fmt.Sprintf("assets error: chapter<%s> was not found in novel<%s>", chapterId, novelId))
	}

	c.String(http.StatusOK, chapterContent)
}

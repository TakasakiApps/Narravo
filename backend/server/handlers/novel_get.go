package handlers

import (
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/dao"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/ohanakogo/exceptiongo"
)

var GetNovelInfo gin.HandlerFunc = func(c *gin.Context) {
	novelId := c.Param("novel_id")

	novel := dao.GetInstance().QueryNovelById(novelId)

	if novel == nil {
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException](fmt.Sprintf("novel<%s> not found", novelId))
	}

	c.JSON(
		200,
		map[string]string{
			"name":        novel.Name,
			"description": novel.Description,
			"authorId":    novel.AuthorId,
			"coverId":     novel.CoverId,
		},
	)
}

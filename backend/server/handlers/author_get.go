package handlers

import (
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/dao"
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/strutil"
	"github.com/ohanakogo/exceptiongo"
	"math"
	"net/http"
)

func toResponsibleAuthorInfo(author *entity.Author) map[string]any {
	return map[string]any{
		"name":        author.Name,
		"description": author.Description,
		"avatarId":    author.AvatarId,
	}
}

var GetAllAuthorInfo gin.HandlerFunc = func(c *gin.Context) {
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

	countAuthor := dao.GetInstance().CountAuthor()
	maxPage := int(math.Ceil(float64(countAuthor) / float64(countInt)))

	var authorInfoList []map[string]any

	if pageInt <= maxPage {
		authors := dao.GetInstance().QueryAuthorsLimit((pageInt-1)*countInt, countInt)
		for _, author := range authors {
			authorInfoList = append(authorInfoList, toResponsibleAuthorInfo(author))
		}
	}

	c.JSON(http.StatusOK, map[string]any{
		"maxPage": maxPage,
		"list":    authorInfoList,
	})
}

var GetAuthorInfo gin.HandlerFunc = func(c *gin.Context) {
	authorId := c.Param("authorId")

	author := dao.GetInstance().QueryAuthorById(authorId)
	if author == nil {
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException](fmt.Sprintf("author<%s> not found", authorId))
	}

	c.JSON(http.StatusOK, toResponsibleAuthorInfo(author))
}
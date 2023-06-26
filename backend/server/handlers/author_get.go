package handlers

import (
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/dao"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/ohanakogo/exceptiongo"
	"net/http"
)

var GetAuthorInfo gin.HandlerFunc = func(c *gin.Context) {
	authorId := c.Param("author_id")

	author := dao.GetInstance().QueryAuthorById(authorId)
	if author == nil {
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException](fmt.Sprintf("author<%s> not found", authorId))
	}

	c.JSON(http.StatusOK, map[string]string{
		"name":        author.Name,
		"description": author.Description,
		"avatarId":    author.AvatarId,
	})
}

package handlers

import (
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/dao"
	"github.com/TakasakiApps/Narravo/backend/internal/assets"
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/ohanakogo/exceptiongo"
	"net/http"
)

func checkPostAuthor(postAuthor *entity.PostAuthor) {
	_, exists := assets.ReadAvatar(postAuthor.AvatarId)
	if !exists {
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException]("avatar not found, you must upload a avatar before add author")
	}
}

var AddAuthor gin.HandlerFunc = func(c *gin.Context) {
	author := &entity.PostAuthor{}
	utils.ConvMapToStructure(utils.GetData(c), author)

	checkPostAuthor(author)

	effected := dao.GetInstance().AddAuthor(&entity.Author{
		ID:         GenerateUniqueID(),
		PostAuthor: *author,
	})

	if effected == 1 {
		c.Status(http.StatusOK)
	} else {
		exceptiongo.QuickThrowMsg[types.ServerNotModifiedException]("add author failed")
	}
}

var UpdateAuthorInfo gin.HandlerFunc = func(c *gin.Context) {
	authorId := c.Param("authorId")

	author := dao.GetInstance().QueryAuthorById(authorId)
	if author == nil {
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException](fmt.Sprintf("author<%s> not found", authorId))
	}

	postAuthor := &entity.PostAuthor{}
	utils.ConvMapToStructure(utils.GetData(c), postAuthor)

	checkPostAuthor(postAuthor)

	effected := dao.GetInstance().UpdateAuthor(&entity.Author{
		ID:         authorId,
		PostAuthor: *postAuthor,
	})

	if effected == 1 {
		c.Status(http.StatusOK)
	} else {
		exceptiongo.QuickThrowMsg[types.ServerNotModifiedException]("update author failed")
	}
}

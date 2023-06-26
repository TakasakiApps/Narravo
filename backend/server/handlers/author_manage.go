package handlers

import (
	"github.com/TakasakiApps/Narravo/backend/dao"
	"github.com/TakasakiApps/Narravo/backend/internal/assets"
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/ohanakogo/exceptiongo"
)

var AddAuthor gin.HandlerFunc = func(c *gin.Context) {
	author := &entity.PostAuthor{}
	utils.ConvMapToStructure(utils.GetData(c), author)

	_, exists := assets.ReadAvatar(author.AvatarId)
	if !exists {
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException]("avatar not found, you must upload a avatar before add author")
	}

	dao.GetInstance().AddAuthor(&entity.Author{
		ID:         GenerateUniqueID(),
		PostAuthor: *author,
	})
}

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
		ID:        utils.GenerateUniqueID(),
		PostNovel: *postNovel,
	})
}

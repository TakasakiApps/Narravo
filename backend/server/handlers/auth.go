package handlers

import (
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/config"
	"github.com/TakasakiApps/Narravo/backend/dao"
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"github.com/TakasakiApps/Narravo/backend/internal/global"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/ohanakogo/exceptiongo"
	"github.com/ohanakogo/exceptiongo/pkg/etype"
	"net/http"
)

var Register gin.HandlerFunc = func(c *gin.Context) {
	user := &entity.User{}
	utils.ConvMapToStructure(utils.GetData(c), user)

	queryUser := dao.GetInstance().QueryUserByName(user)
	if queryUser != nil {
		exceptiongo.QuickThrowMsg[types.ServerUnauthorizedException](fmt.Sprintf("user%v already existed", *user))
	}

	effected := dao.GetInstance().AddUser(user)

	if effected == 1 {
		c.JSON(http.StatusOK, user)
	} else {
		exceptiongo.QuickThrowMsg[types.DataBaseOperationFailedException]("register failed")
	}
}

var Login gin.HandlerFunc = func(c *gin.Context) {
	user := &entity.User{}
	utils.ConvMapToStructure(utils.GetData(c), user)

	queryUser := dao.GetInstance().QueryUserByName(user)
	if queryUser == nil {
		exceptiongo.QuickThrowMsg[types.ServerUnauthorizedException](fmt.Sprintf("user%v is not registered", *user))
	} else if queryUser.Password == user.Password {
		c.JSON(http.StatusOK, entity.UserToken{
			Token: utils.JWTSign[entity.User](*queryUser, config.GetInstance().Crypto.TokenAesKey, global.TokenExpireDelay),
		})
	} else {
		exceptiongo.QuickThrowMsg[types.ServerUnauthorizedException]("incorrect password provided")
	}
}

var Renew gin.HandlerFunc = func(c *gin.Context) {
	userToken := &entity.UserToken{}
	utils.ConvMapToStructure(utils.GetData(c), userToken)

	defer exceptiongo.TryHandle[types.JWTParseFailedException](func(e *etype.Exception) {
		exceptiongo.QuickThrowMsg[types.ServerUnauthorizedException](e.Error())
	})

	utils.JWTParse[entity.User](userToken.Token, config.GetInstance().Crypto.TokenAesKey, func(isValid bool, obj entity.User) {
		if !isValid {
			exceptiongo.QuickThrowMsg[types.ServerUnauthorizedException]("token is invalid")
		}

		user := dao.GetInstance().QueryUserByName(&obj)
		if user == nil {
			exceptiongo.QuickThrowMsg[types.ServerUnauthorizedException]("token contains invalid data")
		}

		if user.Password != obj.Password {
			exceptiongo.QuickThrowMsg[types.ServerUnauthorizedException]("user password has expired")
		}

		c.JSON(http.StatusOK, entity.UserToken{
			Token: utils.JWTSign[entity.User](obj, config.GetInstance().Crypto.TokenAesKey, global.TokenExpireDelay),
		})
	})
}

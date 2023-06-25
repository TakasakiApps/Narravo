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
	"net/http"
)

var Register gin.HandlerFunc = func(c *gin.Context) {
	user := &entity.User{}
	utils.ConvMapToStructure(utils.GetData(c), user)

	queryUser := dao.GetInstance().QueryUserByName(user)
	if queryUser != nil {
		exceptiongo.QuickThrowMsg[types.ServerUnauthorizedException](fmt.Sprintf("user%v already existed", *user))
	}

	// hashing password
	user.Password = utils.MD5(user.Password)

	effected := dao.GetInstance().AddUser(&entity.User{
		Name:     user.Name,
		Password: utils.MD5(user.Password),
	})

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
	} else if queryUser.Password == utils.MD5(user.Password) {
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

	user := utils.JWTParseObj[entity.User](userToken.Token, config.GetInstance().Crypto.TokenAesKey)

	c.JSON(http.StatusOK, entity.UserToken{
		Token: utils.JWTSign[entity.User](user, config.GetInstance().Crypto.TokenAesKey, global.TokenExpireDelay),
	})
}

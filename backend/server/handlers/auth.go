package handlers

import (
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/config"
	"github.com/TakasakiApps/Narravo/backend/dao"
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/ohanakogo/exceptiongo"
	"net/http"
)

var Register gin.HandlerFunc = func(c *gin.Context) {
	user := &entity.User{}
	utils.ConvMapToStructure(utils.GetData(c), user)

	queryUser := dao.GetInstance().QueryUser(user.Name)
	if queryUser == nil {
		exceptiongo.QuickThrowMsg[types.RuntimeException](fmt.Sprintf("user%v already existed", *user))
	}

	effected := dao.GetInstance().AddUser(user)

	if effected == 1 {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusBadRequest, nil)
	}
}

var Login gin.HandlerFunc = func(c *gin.Context) {
	user := &entity.User{}
	utils.ConvMapToStructure(utils.GetData(c), user)

	queryUser := dao.GetInstance().QueryUser(user.Name)
	if queryUser.Password == user.Password {
		c.JSON(http.StatusOK, entity.UserToken{
			Token: utils.JWTSign[entity.User](*queryUser, config.GetInstance().Crypto.AesKey, 1296000),
		})
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

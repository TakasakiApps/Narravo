package handlers

import (
	"fmt"
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

	queryUser := dao.Get().QueryUser(user.Name)
	if queryUser == nil {
		exceptiongo.QuickThrowMsg[types.RuntimeException](fmt.Sprintf("user%v already existed", *user))
	}

	effected := dao.Get().AddUser(user)

	if effected == 1 {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusBadRequest, nil)
	}
}

var Login gin.HandlerFunc = func(c *gin.Context) {
	user := &entity.User{}
	utils.ConvMapToStructure(utils.GetData(c), user)

	queryUser := dao.Get().QueryUser(user.Name)
	if queryUser.Password == user.Password {
		c.JSON(http.StatusOK, queryUser)
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

package handlers

import (
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/ohanakogo/exceptiongo"
)

var NovelIdCheckerComponent gin.HandlerFunc = func(c *gin.Context) {
	novelId := c.Param("novelId")

	if novelId == "" {
		c.Abort()
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException]("novelId can't be empty")
	}
}

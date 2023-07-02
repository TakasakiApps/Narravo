package handlers

import (
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/strutil"
	"github.com/ohanakogo/exceptiongo"
)

var NovelIdCheckerComponent gin.HandlerFunc = func(c *gin.Context) {
	if !strutil.IsEndOf(c.FullPath(), "allInfo") {
		novelId := c.Param("novelId")

		if novelId == "" {
			c.Abort()
			exceptiongo.QuickThrowMsg[types.ServerBadRequestException]("novelId can't be empty")
		}

		c.Next()
	}
}

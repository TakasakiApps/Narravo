package handlers

import (
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/strutil"
	"github.com/ohanakogo/exceptiongo"
)

var AuthorIdCheckerComponent gin.HandlerFunc = func(c *gin.Context) {
	if !strutil.IsEndOf(c.FullPath(), "allInfo") {
		authorId := c.Param("authorId")

		if authorId == "" {
			c.Abort()
			exceptiongo.QuickThrowMsg[types.ServerBadRequestException]("authorId can't be empty")
		}

		c.Next()
	}
}

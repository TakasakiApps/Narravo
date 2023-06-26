package handlers

import (
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/ohanakogo/exceptiongo"
)

var AuthorIdCheckerComponent gin.HandlerFunc = func(c *gin.Context) {
	authorId := c.Param("authorId")

	if authorId == "" {
		c.Abort()
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException]("authorId can't be empty")
	}
}

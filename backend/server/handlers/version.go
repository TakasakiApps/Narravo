package handlers

import (
	"github.com/TakasakiApps/Narravo/backend/internal/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

var GetVersion gin.HandlerFunc = func(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]any{
		"version": global.GetVersion(),
	})
}

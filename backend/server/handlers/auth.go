package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Auth gin.HandlerFunc = func(c *gin.Context) {
	// Test response
	c.JSON(http.StatusOK, struct {
		OK bool
	}{true})
}

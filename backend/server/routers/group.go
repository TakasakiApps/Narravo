package routers

import (
	"github.com/TakasakiApps/Narravo/backend/server/engine"
	"github.com/gin-gonic/gin"
)

var api *gin.RouterGroup

var v1 *gin.RouterGroup

func registerRouterGroup() {
	api = engine.Gin.Group("/api")
	v1 = api.Group("/v1")
}

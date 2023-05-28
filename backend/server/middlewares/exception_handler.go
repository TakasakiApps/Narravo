package middlewares

import (
	"github.com/TakasakiApps/Narravo/backend/internal/global"
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/ohanakogo/exceptiongo"
	"github.com/ohanakogo/exceptiongo/pkg/etype"
	"net/http"
)

func exceptionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer exceptiongo.NewExceptionHandler(func(exception *etype.Exception) {
			switch exception.Type() {
			default:
				global.GetLogger().Error(exception.GetStackTraceMessage())
				utils.NewContextUtil(c).JsonException(http.StatusBadRequest, exception)
			}
		}).Deploy()

		// calls the request handler trigger to make the defer exception handler available.
		c.Next()
	}
}

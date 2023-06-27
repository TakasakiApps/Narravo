package middlewares

import (
	"github.com/TakasakiApps/Narravo/backend/internal/global"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/ohanakogo/exceptiongo"
	"github.com/ohanakogo/exceptiongo/pkg/etype"
	"github.com/ohanakogo/ohanakoutilgo"
	"net/http"
)

func getExceptionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer exceptiongo.NewExceptionHandler(func(exception *etype.Exception) {
			global.GetLogger().Error(exception.GetStackTraceMessage())
			contextUtil := utils.NewContextUtil(c)
			switch exception.Type() {
			case ohanakoutilgo.TypeOf[types.ServerUnauthorizedException]():
				contextUtil.JsonException(http.StatusUnauthorized, exception)
			case ohanakoutilgo.TypeOf[types.ServerNotModifiedException]():
				contextUtil.JsonException(http.StatusBadRequest, exception)
			default:
				contextUtil.JsonException(http.StatusBadRequest, exception)
			}
		}).Deploy()

		// calls the request handler trigger to make the defer exception handler available.
		c.Next()
	}
}

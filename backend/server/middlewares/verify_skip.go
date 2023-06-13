//go:build !verify

package middlewares

import (
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/strutil"
	"net/http"
)

func getVerificationComponent() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := make(map[string]any)
		switch c.Request.Method {
		case http.MethodPost:
			_ = c.ShouldBindJSON(&data)
		case http.MethodGet:
			for k, v := range c.Request.URL.Query() {
				data[k] = strutil.Join("|", v...)
			}
		}
		utils.PutData(c, data)
		// SKIP
		c.Next()
	}
}

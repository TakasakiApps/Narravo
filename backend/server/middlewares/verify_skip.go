//go:build !verify

package middlewares

import (
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/strutil"
	"net/http"
	"strings"
)

func getVerificationComponent() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := make(map[string]any)
		switch {
		case c.Request.Method == http.MethodGet || strings.Contains(c.FullPath(), "/assets/upload"):
			for k, v := range c.Request.URL.Query() {
				data[k] = strutil.Join("|", v...)
			}
		case c.Request.Method == http.MethodPost:
			_ = c.ShouldBindJSON(&data)
		}
		utils.PutData(c, data)
		// SKIP
		c.Next()
	}
}

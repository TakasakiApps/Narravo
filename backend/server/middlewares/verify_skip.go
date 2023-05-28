//go:build !verify

package middlewares

import "github.com/gin-gonic/gin"

func getVerificationComponent() gin.HandlerFunc {
	return func(c *gin.Context) {
		// SKIP
		c.Next()
	}
}

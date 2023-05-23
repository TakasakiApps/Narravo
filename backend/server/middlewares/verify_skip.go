//go:build !verify

package middlewares

import "github.com/gin-gonic/gin"

func verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		// SKIP
		c.Next()
	}
}

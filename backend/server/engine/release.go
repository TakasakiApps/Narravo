//go:build !release

package engine

import "github.com/gin-gonic/gin"

func init() {
	gin.SetMode("release")
}

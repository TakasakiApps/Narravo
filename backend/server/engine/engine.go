package engine

import (
	"github.com/gin-gonic/gin"
)

var Gin *gin.Engine

func init() {
	Gin = gin.Default()
}

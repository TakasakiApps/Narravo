package engine

import (
	"github.com/gin-gonic/gin"
)

var Gin *gin.Engine

func setGin() {
	setMode()
}

func init() {
	setGin()
	Gin = gin.Default()
}

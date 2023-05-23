package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/ohanakogo/exceptiongo/pkg/etype"
)

type ContextUtil struct {
	*gin.Context
}

func NewContextUtil(context *gin.Context) *ContextUtil {
	return &ContextUtil{context}
}

func (c *ContextUtil) JsonError(code int, errMsg string) {
	c.JSON(code, map[string]any{
		"error": errMsg,
	})
}

func (c *ContextUtil) JsonException(code int, exception *etype.Exception) {
	c.JSON(code, map[string]any{
		"error": exception.Error(),
	})
}

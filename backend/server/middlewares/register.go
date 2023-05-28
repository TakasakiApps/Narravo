package middlewares

import (
	"github.com/TakasakiApps/Narravo/backend/server/engine"
)

func Register() {
	engine.Gin.Use(exceptionHandler())
	engine.Gin.Use(verify())
}

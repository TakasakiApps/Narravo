package middlewares

import (
	"github.com/TakasakiApps/Narravo/backend/server/engine"
)

func Register() {
	engine.Gin.Use(verify())
}

package middlewares

import (
	"github.com/TakasakiApps/Narravo/backend/server/engine"
)

func Register() {
	engine.Gin.Use(getExceptionHandler())
	engine.Gin.Use(getVerificationComponent())
	engine.Gin.Use(getUserAuthorization("/register", "/login", "/reset/password"))
}

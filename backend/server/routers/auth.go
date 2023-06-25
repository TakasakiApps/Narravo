package routers

import (
	"github.com/TakasakiApps/Narravo/backend/server/handlers"
)

func registerAuthApi() {
	authGroup := v1.Group("/auth")
	authGroup.POST("/register", handlers.Register)
	authGroup.POST("/login", handlers.LegacyAuthorizationComponent, handlers.Login)
	authGroup.POST("/renew", handlers.Renew)

	resetGroup := authGroup.Group("/reset")
	resetGroup.POST("/password", handlers.LegacyAuthorizationComponent, handlers.ResetPassword)
}

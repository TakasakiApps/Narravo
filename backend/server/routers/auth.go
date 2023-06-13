package routers

import (
	"github.com/TakasakiApps/Narravo/backend/server/handlers"
)

func registerAuthApi() {
	authGroup := v1.Group("/auth")
	authGroup.POST("/register", handlers.Register)
	authGroup.POST("/login", handlers.Login)
}

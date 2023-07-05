package routers

import "github.com/TakasakiApps/Narravo/backend/server/handlers"

func registerSensitiveApi() {
	sensitiveGroup := v1.Group("/sensitive")
	sensitiveGroup.POST("check", handlers.SensitiveCheck)
}

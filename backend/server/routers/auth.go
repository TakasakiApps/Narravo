package routers

import (
	"github.com/TakasakiApps/Narravo/backend/server/handlers"
)

func registerAuthApi() {
	v1.POST("/auth", handlers.Auth)
}

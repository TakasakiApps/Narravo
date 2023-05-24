package routers

import "github.com/TakasakiApps/Narravo/backend/server/handlers"

func registerVersionApi() {
	v1.GET("/version", handlers.GetVersion)
}

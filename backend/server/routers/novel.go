package routers

import "github.com/TakasakiApps/Narravo/backend/server/handlers"

func registerNovelApi() {
	novelGroup := v1.Group("/novel")

	getNovelGroup := novelGroup.Group("/get/:novel_id")
	getNovelGroup.GET("/info", handlers.GetNovelInfo)
}

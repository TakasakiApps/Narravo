package routers

import "github.com/TakasakiApps/Narravo/backend/server/handlers"

func registerNovelApi() {
	novelGroup := v1.Group("/novel")

	novelGroup.POST("/add", handlers.CheckAdminPermissionComponent, handlers.AddNovel)

	getGroup := novelGroup.Group("/get/:novel_id")
	getGroup.GET("/info", handlers.GetNovelInfo)
	getGroup.GET("/catalog", handlers.GetNovelCatalog)
	getGroup.GET("/chapter", handlers.GetNovelChapter)
}

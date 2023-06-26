package routers

import "github.com/TakasakiApps/Narravo/backend/server/handlers"

func registerNovelApi() {
	novelGroup := v1.Group("/novel")

	novelGroup.POST("/add", handlers.CheckAdminPermissionComponent, handlers.AddNovel)

	getNovelGroup := novelGroup.Group("/get/:novel_id")
	getNovelGroup.GET("/info", handlers.GetNovelInfo)
	getNovelGroup.GET("/catalog", handlers.GetNovelCatalog)
	getNovelGroup.GET("/chapter", handlers.GetNovelChapter)
}

package routers

import "github.com/TakasakiApps/Narravo/backend/server/handlers"

func registerNovelApi() {
	novelGroup := v1.Group("/novel")

	novelGroup.POST(
		"/add",
		handlers.CheckAdminPermissionComponent,
		handlers.AddNovel,
	)
	novelGroup.GET("/search", handlers.SearchNovelInfo)

	deleteGroup := novelGroup.Group(
		"/delete/:novelId",
		handlers.NovelIdCheckerComponent,
	)
	deleteGroup.GET("/", handlers.DeleteNovel)

	getGroup := novelGroup.Group(
		"/get/:novelId",
		handlers.NovelIdCheckerComponent,
	)
	getGroup.GET("/info", handlers.GetNovelInfo)
	getGroup.GET("/allInfo", handlers.GetAllNovelInfo)
	getGroup.GET("/catalog", handlers.GetNovelCatalog)
	getGroup.GET("/chapter", handlers.GetNovelChapter)

	updateGroup := novelGroup.Group(
		"/update/:novelId",
		handlers.NovelIdCheckerComponent,
		handlers.CheckAdminPermissionComponent,
	)
	updateGroup.POST("/info", handlers.UpdateNovelInfo)
	updateGroup.POST("/catalog", handlers.UpdateNovelCatalog)
	updateGroup.POST("/chapter", handlers.UpdateNovelChapter)
}

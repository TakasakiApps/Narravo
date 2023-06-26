package routers

import "github.com/TakasakiApps/Narravo/backend/server/handlers"

func registerAuthorApi() {
	authorGroup := v1.Group("/author")

	authorGroup.POST("/add", handlers.CheckAdminPermissionComponent, handlers.AddAuthor)

	getGroup := authorGroup.Group("/get/:author_id")
	getGroup.GET("/info", handlers.GetAuthorInfo)
}

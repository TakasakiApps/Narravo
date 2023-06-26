package routers

import "github.com/TakasakiApps/Narravo/backend/server/handlers"

func registerAuthorApi() {
	authorGroup := v1.Group("/author")

	authorGroup.POST(
		"/add",
		handlers.CheckAdminPermissionComponent,
		handlers.AddAuthor,
	)

	getGroup := authorGroup.Group(
		"/get/:authorId",
		handlers.AuthorIdCheckerComponent,
	)
	getGroup.GET("/info", handlers.GetAuthorInfo)

	updateGroup := authorGroup.Group(
		"/update/:authorId",
		handlers.AuthorIdCheckerComponent,
		handlers.CheckAdminPermissionComponent,
	)
	updateGroup.POST("/info", handlers.UpdateAuthorInfo)
}

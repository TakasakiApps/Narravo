package routers

import "github.com/TakasakiApps/Narravo/backend/server/handlers"

func registerAuthorApi() {
	authorGroup := v1.Group("/author")

	authorGroup.POST(
		"/add",
		handlers.CheckAdminPermissionComponent,
		handlers.AddAuthor,
	)
	authorGroup.GET("/search", handlers.SearchAuthorInfo)
	authorGroup.GET("/getAllInfo", handlers.GetAllAuthorInfo)

	deleteGroup := authorGroup.Group(
		"delete/:authorId",
		handlers.AuthorIdCheckerComponent,
	)
	deleteGroup.GET("/", handlers.DeleteAuthor)

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

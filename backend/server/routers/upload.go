package routers

import "github.com/TakasakiApps/Narravo/backend/server/handlers"

func registerAssetsApi() {
	assetsGroup := v1.Group("/assets")

	uploadGroup := assetsGroup.Group("/upload")
	fetchGroup := assetsGroup.Group("/fetch")

	uploadGroup.POST(
		"/image",
		handlers.GetUploadCheckerComponent("image/png", "image/jpeg", "image/gif"),
		handlers.UploadImage,
	)
	fetchGroup.GET("/image", handlers.FetchImage)
}

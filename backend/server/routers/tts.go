package routers

import "github.com/TakasakiApps/Narravo/backend/server/handlers"

func registerTTSApi() {
	ttsGroup := v1.Group("/tts", handlers.MustNotGuestComponent)

	ttsGroup.GET("/models", handlers.GetModelInfos)
	ttsGroup.POST("/createTask", handlers.CreateTTSTask)
	ttsGroup.GET("/queryResult", handlers.QueryTTSTaskResult)
}

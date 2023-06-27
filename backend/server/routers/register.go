package routers

func Register() {
	registerRouterGroup()

	registerVersionApi()
	registerAuthApi()
	registerAuthorApi()
	registerNovelApi()
	registerAssetsApi()
}

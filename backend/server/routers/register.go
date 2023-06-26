package routers

func Register() {
	registerRouterGroup()

	registerVersionApi()
	registerAuthApi()
	registerNovelApi()
}

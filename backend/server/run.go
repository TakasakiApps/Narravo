package server

import (
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/TakasakiApps/Narravo/backend/server/engine"
	"github.com/TakasakiApps/Narravo/backend/server/middlewares"
	"github.com/TakasakiApps/Narravo/backend/server/routers"
	"github.com/ohanakogo/exceptiongo"
)

func Registries() {
	middlewares.Register()
	routers.Register()
}

func Run(port uint16) {
	Registries()
	err := engine.Gin.Run(fmt.Sprintf(":%v", port))
	exceptiongo.QuickThrow[types.ServerAbortedException](err)
}

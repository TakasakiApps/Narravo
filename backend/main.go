package main

import (
	"flag"
	"github.com/TakasakiApps/Narravo/backend/bootstrap"
	"github.com/TakasakiApps/Narravo/backend/internal/global"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/TakasakiApps/Narravo/backend/server"
	"github.com/gookit/goutil/fsutil"
	"github.com/ohanakogo/exceptiongo"
	"github.com/ohanakogo/exceptiongo/pkg/etype"
	"github.com/ohanakogo/ohanakoutilgo"
	"math"
)

var (
	Port       uint
	DotEnvFile string
)

func init() {
	flag.UintVar(&Port, "p", 8080, "set the port number for the server")
	flag.StringVar(&DotEnvFile, "e", global.DotEnvFile, "use .env file")

	flag.Parse()

	checkArgs()
}

func checkArgs() {
	// Check if the port number is valid.
	if Port > math.MaxUint16 {
		// If not valid, throw an exception with corresponding error message.
		exceptiongo.QuickThrowMsg[types.ArgValueNotValidException]("the port number must be between 0 and 65535")
	}

	if fsutil.FileExists(DotEnvFile) {
		global.DotEnvFile = DotEnvFile
	}
}
func main() {
	defer exceptiongo.NewExceptionHandler(func(exception *etype.Exception) {
		global.GetLogger().Fatal(exception.GetStackTraceMessage())
	}).Deploy()
	bootstrap.Start()
	server.Activate(ohanakoutilgo.CastToNumber[uint16](Port))
}

package global

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
)

var logger *slog.Logger

func init() {
	loggerHandler := handler.NewConsoleHandler(slog.AllLevels)
	loggerHandler.SetFormatter(slog.NewTextFormatter("[{{datetime}}] [{{channel}}] [{{level}}] {{message}} {{data}} {{extra}}\n"))
	logger = slog.NewWithHandlers(loggerHandler)
}

func GetLogger() *slog.Logger {
	return logger
}

package log

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/context"
)

// Set sets a logger in the context.
func Set(ctx echo.Context, logger *slog.Logger) {
	ctx.Set(context.LoggerKey, logger)
}

// Ctx returns the logger stored in context, or provides the default logger if one is not present.
func Ctx(ctx echo.Context) *slog.Logger {
	if l, ok := ctx.Get(context.LoggerKey).(*slog.Logger); ok {
		return l
	}

	return Default()
}

// Default returns the default logger.
func Default() *slog.Logger {
	return slog.Default()
}

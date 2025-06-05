package provider

import (
	"context"

	"github.com/labstack/echo/v4"
)

type contextKey string

const echoContextKey = contextKey("echoContext")

func NewEchoContext(ctx context.Context, eCtx echo.Context) context.Context {
	return context.WithValue(ctx, echoContextKey, eCtx)
}

func FromEchoContext(ctx context.Context) (echo.Context, bool) {
	eCtx, ok := ctx.Value(echoContextKey).(echo.Context)
	return eCtx, ok
}

package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/context"
)

// ServeCachedPage attempts to load a page from the cache by matching on the complete request URL
// If a page is cached for the requested URL, it will be served here and the request terminated.
// Any request made by an authenticated user or that is not a GET will be skipped.
func ServeCachedPage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		// Skip non GET requests
		if ctx.Request().Method != http.MethodGet {
			return next(ctx)
		}

		// Skip if the user is authenticated
		if ctx.Get(context.AuthenticatedUserKey) != nil {
			return next(ctx)
		}

		// TODO keep this functionality?
		return next(ctx)
	}
}

// CacheControl sets a Cache-Control header with a given max age
func CacheControl(maxAge time.Duration) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			v := "no-cache, no-store"
			if maxAge > 0 {
				v = fmt.Sprintf("public, max-age=%.0f", maxAge.Seconds())
			}
			ctx.Response().Header().Set("Cache-Control", v)
			return next(ctx)
		}
	}
}

package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/log"
	"github.com/mikestefanello/pagoda/pkg/services"

	libstore "github.com/eko/gocache/lib/v4/store"
	"github.com/labstack/echo/v4"
)

// ServeCachedPage attempts to load a page from the cache by matching on the complete request URL
// If a page is cached for the requested URL, it will be served here and the request terminated.
// Any request made by an authenticated user or that is not a GET will be skipped.
func ServeCachedPage(t *services.TemplateRenderer) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// Skip non GET requests
			if ctx.Request().Method != http.MethodGet {
				return next(ctx)
			}

			// Skip if the user is authenticated
			if ctx.Get(context.AuthenticatedUserKey) != nil {
				return next(ctx)
			}

			// Attempt to load from cache
			page, err := t.GetCachedPage(ctx, ctx.Request().URL.String())

			if err != nil {
				switch {
				case errors.Is(err, &libstore.NotFound{}):
				case context.IsCanceledError(err):
					return nil
				default:
					log.Ctx(ctx).Error("failed getting cached page",
						"error", err,
					)
				}

				return next(ctx)
			}

			// Set any headers
			if page.Headers != nil {
				for k, v := range page.Headers {
					ctx.Response().Header().Set(k, v)
				}
			}

			log.Ctx(ctx).Debug("serving cached page")

			return ctx.HTMLBlob(page.StatusCode, page.HTML)
		}
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
